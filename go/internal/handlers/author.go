package handlers

import (
	"errors"
	"github.com/a-h/templ"
	"gorm.io/gorm"
	"library/internal/store/repo"
	"library/internal/templates"
	errorUtils "library/internal/utils/errors"
	"library/internal/utils/htmx/requestHeaders"
	"library/internal/utils/ui"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type GetAuthorHandler struct {
	authorRepo repo.IAuthorRepo
}

type NewGetAuthorHandlerParams struct {
	AuthorRepo repo.IAuthorRepo
}

func NewGetAuthorHandler(params NewGetAuthorHandlerParams) *GetAuthorHandler {
	return &GetAuthorHandler{
		authorRepo: params.AuthorRepo,
	}
}

func (h *GetAuthorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authorIDPathParam := r.PathValue("author_id")
	authorID, err := strconv.ParseInt(authorIDPathParam, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	authorWithBooks, err := h.authorRepo.GetAuthorWithBooks(r.Context(), authorID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.NotFound(w, r)
		} else {
			errorUtils.ServerError(r.Context(), w, err, "Error getting author")
		}
		return
	}

	c := templates.Author(authorWithBooks)

	hxBoostedHeader := r.Header.Get(requestHeaders.HxBoosted)
	if hxBoostedHeader == "true" {
		originURL, _ := url.Parse(r.Header.Get(requestHeaders.HxCurrentURL))
		var oobSwaps []templ.Component
		if !strings.HasPrefix(originURL.Path, "/books") {
			anchor := ui.PathToAnchor[originURL.Path]
			oobSwaps = []templ.Component{
				templates.EnabledNavbarLink(anchor.Id, anchor.Text, originURL.Path, true),
			}
		}

		err = templates.ContentsWithTitle(c, ui.TitleDefault(authorWithBooks.Author.DisplayName), oobSwaps).Render(r.Context(), w)
	} else {
		err = templates.Layout(c, ui.TitleDefault(authorWithBooks.Author.DisplayName), r.URL.Path).Render(r.Context(), w)
	}

	if err != nil {
		errorUtils.ServerError(r.Context(), w, err, "Error rendering book", slog.Int64("authorID", authorID))
	}
}
