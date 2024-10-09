package handlers

import (
	"library/internal/templates"
	"library/internal/utils/errors"
	"library/internal/utils/htmx/requestHeaders"
	"log/slog"
	"net/http"
	"time"
)

type LogoutHandler struct {
	sessionCookieName string
}

type LogoutHandlerParams struct {
	SessionCookieName string
}

func NewLogoutHandler(params LogoutHandlerParams) *LogoutHandler {
	return &LogoutHandler{
		sessionCookieName: params.SessionCookieName,
	}
}

func (h *LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:    h.sessionCookieName,
		MaxAge:  -1,
		Expires: time.Now().Add(-100 * time.Hour),
		Path:    "/",
	})

	hxRequestHeader := r.Header.Get(requestHeaders.HxRequest)
	if hxRequestHeader != "true" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	currentHref := r.URL.Path
	slog.DebugContext(r.Context(), "CurrentHref: "+currentHref)

	if err := templates.SignIn(currentHref).Render(r.Context(), w); err != nil {
		errors.ServerError(r.Context(), w, err, "Error rendering template")
	}
}
