package handlers

import (
	"library/internal/store/repo"
	"library/internal/templates"
	"library/internal/utils/errors"
	"library/internal/utils/htmx/requestHeaders"
	"library/internal/utils/ui"
	"net/http"
	"net/url"
	"strings"

	"github.com/a-h/templ"
)

type GetRegisterHandler struct{}

var _ http.Handler = GetRegisterHandler{}

func NewGetRegisterHandler() GetRegisterHandler {
	return GetRegisterHandler{}
}

func (h GetRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.RegisterPage()
	hxBoostedHeader := r.Header.Get(requestHeaders.HxBoosted)

	var err error
	if hxBoostedHeader != "true" {
		err = templates.Layout(c, ui.TitleRegister, "/register").Render(r.Context(), w)
		if err != nil {
			errors.ServerError(r.Context(), w, err, "Error rendering template")
		}
		return
	}

	originUrl, _ := url.Parse(r.Header.Get(requestHeaders.HxCurrentURL))

	oobSwaps := []templ.Component{
		templates.DisabledNavbarLink(ui.IdAnchorRegister, ui.TextAnchorRegister, true),
	}
	if anchor, anchorExists := ui.PathToAnchor[originUrl.Path]; anchorExists {
		oobSwaps = append(oobSwaps, templates.EnabledNavbarLink(anchor.Id, anchor.Text, originUrl.Path, true))
	}

	err = templates.ContentsWithTitle(c, ui.TitleRegister, oobSwaps).Render(r.Context(), w)

	if err != nil {
		errors.ServerError(r.Context(), w, err, "Error rendering template")
	}
}

type PostRegisterHandler struct {
	userStore repo.IUserRepo
}

var _ http.Handler = &PostRegisterHandler{}

type PostRegisterHandlerParams struct {
	UserRepo repo.IUserRepo
}

func NewPostRegisterHandler(params PostRegisterHandlerParams) *PostRegisterHandler {
	return &PostRegisterHandler{
		userStore: params.UserRepo,
	}
}

func (h *PostRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	firstName := strings.TrimSpace(r.FormValue("first_name"))
	lastName := strings.TrimSpace(r.FormValue("last_name"))
	email := strings.TrimSpace(r.FormValue("email"))
	password := strings.TrimSpace(r.FormValue("password"))

	err := h.userStore.CreateUser(r.Context(), firstName, lastName, email, password)

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
	err = templates.ContentsWithTitle(c, ui.TitleRegister, nil).Render(r.Context(), w)

	if err != nil {
		http.Error(w, "error rendering template", http.StatusInternalServerError)
		return
	}
}
