package handlers

import (
	b64 "encoding/base64"
	"fmt"
	"library/internal/hash"
	"library/internal/store"
	"library/internal/templates"
	"net/http"
	"time"
)

type PostLoginHandler struct {
	userStore         store.UserRepo
	sessionStore      store.SessionRepo
	passwordhash      hash.PasswordHasher
	sessionCookieName string
}

type PostLoginHandlerParams struct {
	UserStore         store.UserRepo
	SessionStore      store.SessionRepo
	PasswordHash      hash.PasswordHasher
	SessionCookieName string
}

func NewPostLoginHandler(params PostLoginHandlerParams) *PostLoginHandler {
	return &PostLoginHandler{
		userStore:         params.UserStore,
		sessionStore:      params.SessionStore,
		passwordhash:      params.PasswordHash,
		sessionCookieName: params.SessionCookieName,
	}
}

func (h *PostLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := h.userStore.GetUser(r.Context(), email)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		c := templates.LoginError()
		c.Render(r.Context(), w)
		return
	}

	passwordIsValid, err := h.passwordhash.ComparePasswordAndHash(password, user.PasswordHash)

	if err != nil || !passwordIsValid {
		w.WriteHeader(http.StatusUnauthorized)
		c := templates.LoginError()
		c.Render(r.Context(), w)
		return
	}

	session, err := h.sessionStore.CreateSession(r.Context(), user.ID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID := user.ID
	sessionID := session.ID

	cookieValue := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%d", sessionID, userID)))

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

	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusOK)
}
