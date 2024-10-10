package handlers

import (
	"errors"
	"github.com/a-h/templ"
	"gorm.io/gorm"
	"library/internal/store/model"
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

type GetBookHandler struct {
	bookRepo repo.IBookRepo
}

type NewGetBookHandlerParams struct {
	BookRepo repo.IBookRepo
}

func NewGetBookHandler(params NewGetBookHandlerParams) *GetBookHandler {
	return &GetBookHandler{
		bookRepo: params.BookRepo,
	}
}

func (h *GetBookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bookIdPathParam := r.PathValue("book_id")
	bookID, err := strconv.ParseInt(bookIdPathParam, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	bookWithAuthors, err := h.bookRepo.GetBookWithAuthors(r.Context(), bookID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.NotFound(w, r)
		} else {
			errorUtils.ServerError(r.Context(), w, err, "Error getting book")
		}
		return
	}

	c := templates.Book(bookWithAuthors)

	hxBoostedHeader := r.Header.Get(requestHeaders.HxBoosted)
	if hxBoostedHeader == "true" {
		originURL, _ := url.Parse(r.Header.Get(requestHeaders.HxCurrentURL))
		var oobSwaps []templ.Component
		if !strings.HasPrefix(originURL.Path, "/authors") {
			anchor := ui.PathToAnchor[originURL.Path]
			oobSwaps = []templ.Component{
				templates.EnabledNavbarLink(anchor.Id, anchor.Text, originURL.Path, true),
			}
		}

		err = templates.ContentsWithTitle(c, buildTitle(bookWithAuthors), oobSwaps).Render(r.Context(), w)
	} else {
		err = templates.Layout(c, buildTitle(bookWithAuthors), r.URL.Path).Render(r.Context(), w)
	}

	if err != nil {
		errorUtils.ServerError(r.Context(), w, err, "Error rendering book", slog.Int64("bookID", bookID))
	}
}

func buildTitle(book model.BookWithAuthors) string {
	var titleSB strings.Builder

	length := 10 + len(book.Book.Title)
	if len(book.Authors) > 0 {
		length += 3 + (len(book.Authors)-1)*2
		for _, author := range book.Authors {
			length += len(author.DisplayName)
		}
	}
	titleSB.Grow(length)

	titleSB.WriteString("Library - ")
	titleSB.WriteString(book.Book.Title)
	if len(book.Authors) > 0 {
		titleSB.WriteString(" - ")
		for i, author := range book.Authors {
			titleSB.WriteString(author.DisplayName)
			if i < len(book.Authors)-1 {
				titleSB.WriteString(", ")
			}
		}
	}

	return titleSB.String()
}
