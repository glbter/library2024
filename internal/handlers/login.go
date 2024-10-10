package handlers

import (
	"library/internal/hash"
	"library/internal/store/repo"
	"library/internal/templates"
	"library/internal/utils/encoders"
	"library/internal/utils/errors"
	"library/internal/utils/htmx/requestHeaders"
	"library/internal/utils/htmx/responseHeaders"
	"library/internal/utils/ui"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/a-h/templ"
)

type GetLoginHandler struct{}

func NewGetLoginHandler() *GetLoginHandler {
	return &GetLoginHandler{}
}

func (h *GetLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.Login()
	hxBoostedHeader := r.Header.Get(requestHeaders.HxBoosted)

	var err error
	if hxBoostedHeader == "true" {
		originUrl, _ := url.Parse(r.Header.Get(requestHeaders.HxCurrentURL))

		oobSwaps := []templ.Component{
			templates.DisabledNavbarLink(ui.IdAnchorLogin, ui.TextAnchorLogin, true),
		}
		if anchor, anchorExists := ui.PathToAnchor[originUrl.Path]; anchorExists {
			oobSwaps = append(oobSwaps, templates.EnabledNavbarLink(anchor.Id, anchor.Text, originUrl.Path, true))
		}

		err = templates.ContentsWithTitle(c, ui.TitleLogin, oobSwaps).Render(r.Context(), w)
	} else {
		err = templates.Layout(c, ui.TitleLogin, "/login").Render(r.Context(), w)
	}

	if err != nil {
		errors.ServerError(r.Context(), w, err, "Error rendering template")
	}
}

type PostLoginHandler struct {
	userStore         repo.IUserRepo
	sessionStore      repo.ISessionRepo
	passwordHasher    hash.PasswordHasher
	sessionCookieName string
}

type PostLoginHandlerParams struct {
	UserStore         repo.IUserRepo
	SessionRepo       repo.ISessionRepo
	PasswordHasher    hash.PasswordHasher
	SessionCookieName string
}

func NewPostLoginHandler(params PostLoginHandlerParams) *PostLoginHandler {
	return &PostLoginHandler{
		userStore:         params.UserStore,
		sessionStore:      params.SessionRepo,
		passwordHasher:    params.PasswordHasher,
		sessionCookieName: params.SessionCookieName,
	}
}

func (h *PostLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := h.userStore.GetUser(r.Context(), email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		if err = templates.LoginError().Render(r.Context(), w); err != nil {
			slog.ErrorContext(r.Context(), "Error rendering login error", slog.Any("err", err))
		}
		return
	}

	passwordIsValid, err := h.passwordHasher.ComparePasswordAndHash(password, user.PasswordHash)
	if err != nil || !passwordIsValid {
		w.WriteHeader(http.StatusUnauthorized)
		if err = templates.LoginError().Render(r.Context(), w); err != nil {
			slog.ErrorContext(r.Context(), "Error rendering login error", slog.Any("err", err))
		}
		return
	}

	session, err := h.sessionStore.CreateSession(r.Context(), user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Error creating session", slog.Any("err", err))
		return
	}

	userID := user.ID
	sessionID := session.ID

	cookieValue := encoders.EncodeCookieValue(sessionID, userID)

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{
		Name:     h.sessionCookieName,
		Value:    cookieValue,
		Expires:  expiration,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, &cookie)

	if r.Header.Get(requestHeaders.HxRequest) == "true" {
		w.Header().Set(responseHeaders.HxRedirect, "/")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
