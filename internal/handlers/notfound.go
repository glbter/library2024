package handlers

import (
	"library/internal/templates"
	"library/internal/utils/ui"
	"log/slog"
	"net/http"
	"strings"
)

type NotFoundHandler struct{}

func NewNotFoundHandler() *NotFoundHandler {
	return &NotFoundHandler{}
}

func (h *NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if acceptHeaders := r.Header.Values("Accept"); len(acceptHeaders) <= 0 || !strings.Contains(acceptHeaders[0], "text/html") {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	err := templates.Layout(templates.NotFound(), ui.TitleNotFound, "").Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Error rendering template", slog.Any("err", err))
	}
}
