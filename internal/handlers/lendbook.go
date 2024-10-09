package handlers

import (
	"library/internal/middleware"
	"library/internal/store/repo"
	"log/slog"
	"net/http"
	"strconv"
)

type LendBookHandler struct {
	bookRepo repo.IBookRepo
}

type NewLendBookHandlerParams struct {
	BookRepo repo.IBookRepo
}

func NewLendBookHandler(params NewLendBookHandlerParams) *LendBookHandler {
	return &LendBookHandler{
		bookRepo: params.BookRepo,
	}
}

func (h *LendBookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	bookIdFormParams, ok := r.Form["book_id"]
	if !ok || len(bookIdFormParams) == 0 {
		http.Error(w, "book_id required", http.StatusBadRequest)
		return
	}

	bookID, err := strconv.Atoi(bookIdFormParams[0])
	if err != nil {
		http.Error(w, "book_id must be int", http.StatusBadRequest)
		return
	}

	user := middleware.GetUser(ctx)
	if user == nil {
		http.Error(w, "No user in the context", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Extract user from context")
		return
	}

	if err = h.bookRepo.RequestBook(ctx, user.ID, int64(bookID)); err != nil {
		http.Error(w, "Lending a book", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Error lending a book", slog.Any("err", err))
		return
	}

	//c := templates.Index(books)
	//
	//err = templates.Layout(c, "Library").Render(r.Context(), w)
	//
	//if err != nil {
	//	http.Error(w, "Error rendering template", http.StatusInternalServerError)
	//	slog.ErrorContext(r.Context(), "Error rendering template", slog.Any("err", err))
	//}
}
