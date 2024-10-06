package handlers

import (
	"library/internal/templates"
	"log/slog"
	"net/http"
)

type GetRegisterHandler struct{}

func NewGetRegisterHandler() GetRegisterHandler {
	return GetRegisterHandler{}
}

func (h GetRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.RegisterPage()
	err := templates.Layout(c, "Library - Register").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Error rendering template", slog.Any("err", err))
	}
}
