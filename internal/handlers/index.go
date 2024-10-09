package handlers

import (
	"library/internal/store/repo"
	"library/internal/templates"
	"library/internal/utils/errors"
	"library/internal/utils/htmx/requestHeaders"
	"library/internal/utils/ui"
	"net/http"
	"net/url"
	"strconv"

	"github.com/a-h/templ"
)

type IndexHandler struct {
	bookRepo repo.IBookRepo
}

type NewIndexHandlerParams struct {
	BookRepo repo.IBookRepo
}

func NewIndexHandler(params NewIndexHandlerParams) *IndexHandler {
	return &IndexHandler{
		bookRepo: params.BookRepo,
	}
}

const (
	PageQueryParam  = "page"
	LimitQueryParam = "limit"
	DefaultPage     = uint(0)
	DefaultLimit    = uint(10)
)

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	page := DefaultPage
	pageParamValues, hasPage := q[PageQueryParam]
	if hasPage {
		_page, err := strconv.ParseUint(pageParamValues[0], 10, strconv.IntSize)
		if err == nil {
			page = uint(_page)
		}
	}

	limit := DefaultLimit
	limitParamValues, hasLimit := q[LimitQueryParam]
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

	hxRequestHeader := r.Header.Get(requestHeaders.HxRequest)
	if hxRequestHeader != "true" {
		contents := templates.Index(books, page, totalPages)
		if err = templates.Layout(contents, ui.TitleHome, "/").Render(r.Context(), w); err != nil {
			errors.ServerError(r.Context(), w, err, "Error rendering template")
		}
		return
	}

	hxBoostedHeader := r.Header.Get(requestHeaders.HxBoosted)
	if hxBoostedHeader == "true" {
		contents := templates.Index(books, page, totalPages)

		originUrl, _ := url.Parse(r.Header.Get(requestHeaders.HxCurrentURL))

		oobSwaps := []templ.Component{
			templates.DisabledNavbarLink(ui.IdAnchorHome, ui.TextAnchorHome, true),
		}
		if anchor, anchorExists := ui.PathToAnchor[originUrl.Path]; anchorExists {
			oobSwaps = append(oobSwaps, templates.EnabledNavbarLink(anchor.Id, anchor.Text, originUrl.Path, true))
		}

		err = templates.ContentsWithTitle(contents, ui.TitleHome, oobSwaps).Render(r.Context(), w)
		if err != nil {
			errors.ServerError(r.Context(), w, err, "Error rendering template")
		}
		return
	}

	hxTriggerHeader := r.Header.Get(requestHeaders.HxTrigger)
	showPagination := templates.ShowPagination{
		Top:    hxTriggerHeader == ui.IdPaginationTop,
		Bottom: hxTriggerHeader == ui.IdPaginationBottom,
	}
	if err = templates.BooksListItems(books, page, totalPages, showPagination).Render(r.Context(), w); err != nil {
		errors.ServerError(r.Context(), w, err, "Error rendering template")
	}

}
