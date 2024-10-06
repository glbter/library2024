package handlers

import (
	"library/internal/store"
	"library/internal/templates"
	"log/slog"
	"net/http"
	"strconv"
)

type IndexHandler struct {
	bookRepo store.BookRepo
}

type NewIndexHandlerParams struct {
	BookRepo store.BookRepo
}

func NewIndexHandler(params NewIndexHandlerParams) *IndexHandler {
	return &IndexHandler{
		bookRepo: params.BookRepo,
	}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	page := 1
	pageParamValues, hasPage := q["page"]
	if hasPage {
		var err error
		page, err = strconv.Atoi(pageParamValues[0])
		if err != nil {
			page = 1
		}
	}

	limit := 10
	limitParamValues, hasLimit := q["limit"]
	if hasLimit {
		var err error
		limit, err = strconv.Atoi(limitParamValues[0])
		if err != nil {
			limit = 10
		}
	}

	books, err := h.bookRepo.GetBooksWithAuthors(r.Context(), page, limit)
	if err != nil {
		http.Error(w, "Error getting book list", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Error getting book list", slog.Any("err", err))
		return
	}
	c := templates.Index(books)

	err = templates.Layout(c, "Library").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Error rendering template", slog.Any("err", err))
	}
}
