package handlers

import (
	"library/internal/templates"
	"log/slog"
	"net/http"
)

type NotFoundHandler struct{}

func NewNotFoundHandler() *NotFoundHandler {
	return &NotFoundHandler{}
}

func (h *NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.NotFound()
	err := templates.Layout(c, "Not Found").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Error rendering template", slog.Any("err", err))
	}
}
