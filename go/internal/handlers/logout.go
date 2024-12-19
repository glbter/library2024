package handlers

import (
	"errors"
	"library/internal/templates"
	errorUtils "library/internal/utils/errors"
	"library/internal/utils/htmx/requestHeaders"
	"log/slog"
	"net/http"
	"time"
)

type LogoutHandler struct {
	sessionCookieName string
}

var _ http.Handler = &LogoutHandler{}

func NewLogoutHandler(sessionCookieName string) *LogoutHandler {
	if sessionCookieName == "" {
		panic(errors.New("sessionCookieName is required"))
	}
	return &LogoutHandler{sessionCookieName}
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
		errorUtils.ServerError(r.Context(), w, err, "Error rendering template")
	}
}
