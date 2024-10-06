package handlers

import (
	"library/internal/templates"
	"log/slog"
	"net/http"
)

type GetLoginHandler struct{}

func NewGetLoginHandler() *GetLoginHandler {
	return &GetLoginHandler{}
}

func (h *GetLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.Login()
	err := templates.Layout(c, "Library - Login").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Error rendering template", slog.Any("err", err))
	}
}
