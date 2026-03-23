package content

import (
	"bytes"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
)

type Frontmatter struct {
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Slug        string   `yaml:"slug"`
	Order       int      `yaml:"order"`
	Tags        []string `yaml:"tags"`
	Draft       bool     `yaml:"draft"`
}

type Page struct {
	Frontmatter
	// URL path for this page (e.g. /games/minecraft/getting-started)
	Path string
	// Category derived from directory structure (e.g. "games", "self-hosting")
	Category string
	// Subcategory derived from subdirectory (e.g. "minecraft")
	Subcategory string
	// Rendered HTML content
	HTML string
	// Raw markdown content (for search indexing)
	Raw string
	// Source file path relative to content dir
	SourceFile string
}

type Category struct {
	Name          string
	Slug          string
	Subcategories []Subcategory
}

type Subcategory struct {
	Name  string
	Slug  string
	Pages []Page
}

type Store struct {
	Pages      map[string]*Page // keyed by URL path
	Categories []Category       // ordered for navigation
}

var categoryNames = map[string]string{
	"games":        "Games",
	"self-hosting": "Self-Hosting",
	"gamejanitor":  "Game Janitor",
	"hosting":      "Hosting",
}

var categoryOrder = map[string]int{
	"games":        0,
	"self-hosting": 1,
	"gamejanitor":  2,
	"hosting":      3,
}

func Load(contentDir string) (*Store, error) {
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"),
			),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)

	store := &Store{
		Pages: make(map[string]*Page),
	}

	catMap := make(map[string]map[string][]Page)

	err := filepath.WalkDir(contentDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading %s: %w", path, err)
		}

		fm, body, err := parseFrontmatter(data)
		if err != nil {
			return fmt.Errorf("parsing frontmatter in %s: %w", path, err)
		}

		if fm.Draft {
			slog.Debug("skipping draft", "file", path)
			return nil
		}

		relPath, _ := filepath.Rel(contentDir, path)
		parts := strings.Split(filepath.ToSlash(relPath), "/")

		if len(parts) < 2 {
			slog.Warn("skipping file not in a category directory", "file", relPath)
			return nil
		}

		category := parts[0]
		subcategory := ""
		filename := ""

		if len(parts) == 3 {
			subcategory = parts[1]
			filename = strings.TrimSuffix(parts[2], ".md")
		} else {
			filename = strings.TrimSuffix(parts[1], ".md")
		}

		slug := fm.Slug
		if slug == "" {
			slug = filename
		}

		var urlPath string
		if subcategory != "" {
			urlPath = fmt.Sprintf("/%s/%s/%s", category, subcategory, slug)
		} else {
			urlPath = fmt.Sprintf("/%s/%s", category, slug)
		}

		var htmlBuf bytes.Buffer
		if err := md.Convert(body, &htmlBuf); err != nil {
			return fmt.Errorf("rendering markdown for %s: %w", path, err)
		}

		page := Page{
			Frontmatter: fm,
			Path:        urlPath,
			Category:    category,
			Subcategory: subcategory,
			HTML:        htmlBuf.String(),
			Raw:         string(body),
			SourceFile:  relPath,
		}

		if page.Title == "" {
			page.Title = strings.ReplaceAll(filename, "-", " ")
		}

		store.Pages[urlPath] = &page

		if catMap[category] == nil {
			catMap[category] = make(map[string][]Page)
		}
		subKey := subcategory
		if subKey == "" {
			subKey = "_root"
		}
		catMap[category][subKey] = append(catMap[category][subKey], page)

		slog.Debug("loaded page", "path", urlPath, "title", page.Title)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walking content directory: %w", err)
	}

	store.Categories = buildCategories(catMap)

	slog.Info("content loaded", "pages", len(store.Pages), "categories", len(store.Categories))
	return store, nil
}

func parseFrontmatter(data []byte) (Frontmatter, []byte, error) {
	var fm Frontmatter

	content := string(data)
	if !strings.HasPrefix(content, "---\n") {
		return fm, data, nil
	}

	end := strings.Index(content[4:], "\n---\n")
	if end == -1 {
		return fm, data, nil
	}

	fmData := content[4 : 4+end]
	body := []byte(content[4+end+5:])

	if err := yaml.Unmarshal([]byte(fmData), &fm); err != nil {
		return fm, nil, fmt.Errorf("invalid frontmatter YAML: %w", err)
	}

	return fm, body, nil
}

func buildCategories(catMap map[string]map[string][]Page) []Category {
	var categories []Category

	for catSlug, subs := range catMap {
		cat := Category{
			Slug: catSlug,
			Name: categoryNames[catSlug],
		}
		if cat.Name == "" {
			cat.Name = strings.Title(strings.ReplaceAll(catSlug, "-", " "))
		}

		for subSlug, pages := range subs {
			sort.Slice(pages, func(i, j int) bool {
				if pages[i].Order != pages[j].Order {
					return pages[i].Order < pages[j].Order
				}
				return pages[i].Title < pages[j].Title
			})

			subName := strings.Title(strings.ReplaceAll(subSlug, "-", " "))
			if subSlug == "_root" {
				subName = ""
			}

			cat.Subcategories = append(cat.Subcategories, Subcategory{
				Name:  subName,
				Slug:  subSlug,
				Pages: pages,
			})
		}

		sort.Slice(cat.Subcategories, func(i, j int) bool {
			if cat.Subcategories[i].Slug == "_root" {
				return true
			}
			if cat.Subcategories[j].Slug == "_root" {
				return false
			}
			return cat.Subcategories[i].Name < cat.Subcategories[j].Name
		})

		categories = append(categories, cat)
	}

	sort.Slice(categories, func(i, j int) bool {
		oi, ok := categoryOrder[categories[i].Slug]
		if !ok {
			oi = 999
		}
		oj, ok := categoryOrder[categories[j].Slug]
		if !ok {
			oj = 999
		}
		return oi < oj
	})

	return categories
}
