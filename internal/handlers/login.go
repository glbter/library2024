package handlers

import (
	"library/internal/hash"
	"library/internal/store"
	"library/internal/templates"
	"library/internal/utils"
	"log/slog"
	"net/http"
	"time"
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

type PostLoginHandler struct {
	userStore         store.UserRepo
	sessionStore      store.SessionRepo
	passwordHasher    hash.PasswordHasher
	sessionCookieName string
}

type PostLoginHandlerParams struct {
	UserStore         store.UserRepo
	SessionRepo       store.SessionRepo
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
	if !sessionID.Valid {
		w.WriteHeader(http.StatusInternalServerError)
		slog.ErrorContext(r.Context(), "Invalid sessionID", slog.Int64("userID", userID))
		return
	}

	cookieValue := utils.EncodeCookieValue(sessionID.Bytes, userID)

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
