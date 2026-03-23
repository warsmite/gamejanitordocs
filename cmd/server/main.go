package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	"github.com/warsmite/gamejanitordocs/internal/content"
)

type server struct {
	store  *content.Store
	search *content.SearchIndex
	tmpl   *template.Template
	baseURL string
}

type pageData struct {
	Page       *content.Page
	Categories []content.Category
	BaseURL    string
}

type searchData struct {
	Query      string
	Results    []content.SearchResult
	Categories []content.Category
	BaseURL    string
}

type homeData struct {
	Categories []content.Category
	BaseURL    string
}

func main() {
	level := slog.LevelInfo
	if os.Getenv("DEBUG") != "" {
		level = slog.LevelDebug
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level})))

	rootDir := envOr("ROOT_DIR", ".")
	contentDir := filepath.Join(rootDir, "content")
	templateDir := filepath.Join(rootDir, "templates")
	staticDir := filepath.Join(rootDir, "static")
	addr := envOr("ADDR", ":8080")
	baseURL := envOr("BASE_URL", "")

	store, err := content.Load(contentDir)
	if err != nil {
		slog.Error("failed to load content", "error", err)
		os.Exit(1)
	}

	searchIndex, err := content.NewSearchIndex(store)
	if err != nil {
		slog.Error("failed to build search index", "error", err)
		os.Exit(1)
	}
	defer searchIndex.Close()

	funcMap := template.FuncMap{
		"upper": strings.ToUpper,
		"categoryName": func(slug string) string {
			names := map[string]string{
				"games":        "Games",
				"self-hosting": "Self-Hosting",
				"gamejanitor":  "Game Janitor",
				"hosting":      "Hosting",
			}
			if n, ok := names[slug]; ok {
				return n
			}
			return slug
		},
		"safeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
	}

	tmpl, err := template.New("").Funcs(funcMap).ParseGlob(filepath.Join(templateDir, "*.html"))
	if err != nil {
		slog.Error("failed to parse templates", "error", err)
		os.Exit(1)
	}

	srv := &server{
		store:   store,
		search:  searchIndex,
		tmpl:    tmpl,
		baseURL: baseURL,
	}

	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	mux.HandleFunc("GET /search", srv.handleSearch)
	mux.HandleFunc("GET /sitemap.xml", srv.handleSitemap)
	mux.HandleFunc("GET /robots.txt", srv.handleRobots)
	mux.HandleFunc("GET /{$}", srv.handleHome)
	mux.HandleFunc("GET /", srv.handlePage)

	httpServer := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		slog.Info("server starting", "addr", addr, "pages", len(store.Pages))
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server error", "error", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	slog.Info("shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	httpServer.Shutdown(shutdownCtx)
}

func (s *server) handleHome(w http.ResponseWriter, r *http.Request) {
	data := homeData{
		Categories: s.store.Categories,
		BaseURL:    s.baseURL,
	}
	if err := s.tmpl.ExecuteTemplate(w, "home.html", data); err != nil {
		slog.Error("template error", "error", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func (s *server) handlePage(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	page, ok := s.store.Pages[path]
	if !ok {
		http.NotFound(w, r)
		return
	}

	data := pageData{
		Page:       page,
		Categories: s.store.Categories,
		BaseURL:    s.baseURL,
	}
	if err := s.tmpl.ExecuteTemplate(w, "doc.html", data); err != nil {
		slog.Error("template error", "path", path, "error", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func (s *server) handleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	var results []content.SearchResult
	if query != "" {
		var err error
		results, err = s.search.Search(query, 20)
		if err != nil {
			slog.Error("search error", "query", query, "error", err)
		}
	}

	data := searchData{
		Query:      query,
		Results:    results,
		Categories: s.store.Categories,
		BaseURL:    s.baseURL,
	}

	// HTMX partial response — just the results
	if r.Header.Get("HX-Request") == "true" {
		if err := s.tmpl.ExecuteTemplate(w, "search_results", data); err != nil {
			slog.Error("template error", "error", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	if err := s.tmpl.ExecuteTemplate(w, "search.html", data); err != nil {
		slog.Error("template error", "error", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

type sitemapURL struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod,omitempty"`
}

type sitemap struct {
	XMLName xml.Name     `xml:"urlset"`
	XMLNS   string       `xml:"xmlns,attr"`
	URLs    []sitemapURL `xml:"url"`
}

func (s *server) handleSitemap(w http.ResponseWriter, r *http.Request) {
	sm := sitemap{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}

	for path := range s.store.Pages {
		sm.URLs = append(sm.URLs, sitemapURL{
			Loc: s.baseURL + path,
		})
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write([]byte(xml.Header))
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	enc.Encode(sm)
}

func (s *server) handleRobots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "User-agent: *\nAllow: /\nSitemap: %s/sitemap.xml\n", s.baseURL)
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
