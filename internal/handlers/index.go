package handlers

import (
	"library/internal/store"
	"library/internal/templates"
	"library/internal/utils/errors"
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

const (
	DefaultPage  uint = 0
	DefaultLimit uint = 10
)

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	page := DefaultPage
	pageParamValues, hasPage := q["page"]
	if hasPage {
		_page, err := strconv.ParseUint(pageParamValues[0], 10, strconv.IntSize)
		if err == nil {
			page = uint(_page)
		}
	}

	limit := DefaultLimit
	limitParamValues, hasLimit := q["limit"]
	if hasLimit {
		_limit, err := strconv.ParseUint(limitParamValues[0], 10, strconv.IntSize)
		if err == nil {
			limit = uint(_limit)
		}
	}

	books, totalPages, err := h.bookRepo.GetBooksWithAuthors(r.Context(), page, limit)
	if err != nil {
		errors.ServerError(r.Context(), w, err, "Error getting book list")
		return
	}

	if hxRequestHeader := r.Header.Get("HX-Request"); hxRequestHeader == "true" {
		pageStart := r.Header.Get("HX-Trigger") == "books-start" && page >= 1
		showPagination := r.Header.Get("HX-Trigger") == "books-pagination"
		if err = templates.BooksListItems(books, page, totalPages, pageStart, showPagination).Render(r.Context(), w); err != nil {
			errors.ServerError(r.Context(), w, err, "Error rendering template")
		}
		return
	}

	contents := templates.Index(books, page, totalPages)
	if err = templates.Layout(contents, "Library").Render(r.Context(), w); err != nil {
		errors.ServerError(r.Context(), w, err, "Error rendering template")
	}
}
