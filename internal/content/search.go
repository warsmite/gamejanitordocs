package content

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/blevesearch/bleve/v2"
)

type SearchResult struct {
	Path        string
	Title       string
	Description string
	Category    string
	Snippet     string
	Score       float64
}

type SearchIndex struct {
	index bleve.Index
	store *Store
}

type indexDoc struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
	Category    string `json:"category"`
	Tags        string `json:"tags"`
}

func NewSearchIndex(store *Store) (*SearchIndex, error) {
	mapping := bleve.NewIndexMapping()

	docMapping := bleve.NewDocumentMapping()

	textField := bleve.NewTextFieldMapping()
	textField.Analyzer = "en"

	// Boost title matches higher than body
	titleField := bleve.NewTextFieldMapping()
	titleField.Analyzer = "en"
	titleField.Store = true

	docMapping.AddFieldMappingsAt("title", titleField)
	docMapping.AddFieldMappingsAt("description", textField)
	docMapping.AddFieldMappingsAt("body", textField)
	docMapping.AddFieldMappingsAt("category", bleve.NewKeywordFieldMapping())
	docMapping.AddFieldMappingsAt("tags", textField)

	mapping.DefaultMapping = docMapping

	index, err := bleve.NewMemOnly(mapping)
	if err != nil {
		return nil, fmt.Errorf("creating bleve index: %w", err)
	}

	batch := index.NewBatch()
	for path, page := range store.Pages {
		doc := indexDoc{
			Title:       page.Title,
			Description: page.Description,
			Body:        page.Raw,
			Category:    page.Category,
			Tags:        strings.Join(page.Tags, " "),
		}
		batch.Index(path, doc)
	}

	if err := index.Batch(batch); err != nil {
		return nil, fmt.Errorf("indexing content: %w", err)
	}

	slog.Info("search index built", "documents", len(store.Pages))

	return &SearchIndex{index: index, store: store}, nil
}

func (si *SearchIndex) Search(query string, limit int) ([]SearchResult, error) {
	if limit <= 0 {
		limit = 20
	}

	q := bleve.NewQueryStringQuery(query)
	req := bleve.NewSearchRequestOptions(q, limit, 0, false)
	req.Highlight = bleve.NewHighlightWithStyle("html")
	req.Fields = []string{"title", "description", "category"}

	res, err := si.index.Search(req)
	if err != nil {
		return nil, fmt.Errorf("searching: %w", err)
	}

	var results []SearchResult
	for _, hit := range res.Hits {
		page, ok := si.store.Pages[hit.ID]
		if !ok {
			continue
		}

		snippet := ""
		if fragments, ok := hit.Fragments["body"]; ok && len(fragments) > 0 {
			snippet = fragments[0]
		} else if page.Description != "" {
			snippet = page.Description
		}

		results = append(results, SearchResult{
			Path:        hit.ID,
			Title:       page.Title,
			Description: page.Description,
			Category:    page.Category,
			Snippet:     snippet,
			Score:       hit.Score,
		})
	}

	return results, nil
}

func (si *SearchIndex) Close() error {
	return si.index.Close()
}
