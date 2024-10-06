package handlers

import (
	"library/internal/store"
	"library/internal/templates"
	"net/http"
)

type PostRegisterHandler struct {
	userStore store.UserRepo
}

type PostRegisterHandlerParams struct {
	UserRepo store.UserRepo
}

func NewPostRegisterHandler(params PostRegisterHandlerParams) *PostRegisterHandler {
	return &PostRegisterHandler{
		userStore: params.UserRepo,
	}
}

func (h *PostRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	err := h.userStore.CreateUser(r.Context(), email, password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		c := templates.RegisterError()
		err = c.Render(r.Context(), w)
		if err != nil {
			http.Error(w, "error rendering template", http.StatusInternalServerError)
		}
		return
	}

	c := templates.RegisterSuccess()
	err = c.Render(r.Context(), w)

	if err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}
}
