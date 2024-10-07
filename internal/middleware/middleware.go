package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"library/internal/store"
	"library/internal/store/model"
	"library/internal/utils/encoders"
	"log"
	"log/slog"
	"net/http"
	"strings"
)

type key string

var NonceKey key = "nonces"

type Nonces struct {
	Htmx            string
	ResponseTargets string
	Tw              string
	HtmxCSSHashes   []string
}

func generateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func CSPMiddleware(next http.Handler) http.Handler {
	// To use the same nonces in all responses, move the Nonces
	// struct creation to here, outside the handler.

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new Nonces struct for every request when here.
		// move to outside the handler to use the same nonces in all responses
		nonceSet := Nonces{
			Htmx:            generateRandomString(16),
			ResponseTargets: generateRandomString(16),
			Tw:              generateRandomString(16),
			HtmxCSSHashes: []string{
				"sha256-VV50kYRP+CuHcIPnOJ/AV+Q2C0IGVX7AczE6/dxv078=",
				"sha256-cNCcTUjHx0E9D/2VhCEWby6Bd/Ow9sHRp65Rf9s3J68=",
				"sha256-9Ccj/TQB4XGZUjlcvis7DszVKQz1ppCoyRybka1oFIA=",
				"sha256-bsV5JivYxvGywDAZ22EZJKBFip65Ng9xoJVLbBg7bdo=",
			},
		}

		// set nonces in context
		ctx := context.WithValue(r.Context(), NonceKey, nonceSet)
		// insert the nonces into the content security policy header
		cspHeader := fmt.Sprintf("default-src 'self'; script-src 'nonce-%s' 'nonce-%s' ; style-src 'nonce-%s' %s;",
			nonceSet.Htmx,
			nonceSet.ResponseTargets,
			nonceSet.Tw,
			func() string {
				sb := strings.Builder{}
				sb.Grow(len(nonceSet.HtmxCSSHashes)*(51+2) + (len(nonceSet.HtmxCSSHashes) - 1))
				for i, hash := range nonceSet.HtmxCSSHashes {
					sb.WriteByte('\'')
					sb.WriteString(hash)
					sb.WriteByte('\'')
					if i+1 < len(nonceSet.HtmxCSSHashes) {
						sb.WriteByte(' ')
					}
				}
				return sb.String()
			}(),
		)
		w.Header().Set("Content-Security-Policy", cspHeader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TextHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

// GetNonces gets the Nonce from the context, it is a struct called Nonces,
// so we can get the nonce we need by the key, i.e. HtmxNonce
func GetNonces(ctx context.Context) Nonces {
	nonceSet := ctx.Value(NonceKey)
	if nonceSet == nil {
		log.Fatal("error getting nonce set - is nil")
	}

	nonces, ok := nonceSet.(Nonces)
	if !ok {
		log.Fatal("error getting nonce set - not ok")
	}

	return nonces
}

func GetHtmxNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)

	return nonceSet.Htmx
}

func GetResponseTargetsNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)
	return nonceSet.ResponseTargets
}

func GetTwNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)
	return nonceSet.Tw
}

type AuthMiddleware struct {
	sessionStore      store.SessionRepo
	sessionCookieName string
}

func NewAuthMiddleware(sessionStore store.SessionRepo, sessionCookieName string) *AuthMiddleware {
	return &AuthMiddleware{
		sessionStore:      sessionStore,
		sessionCookieName: sessionCookieName,
	}
}

type UserContextKey string

const UserKey UserContextKey = "user"

func (m *AuthMiddleware) AddUserToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sessionCookie, err := r.Cookie(m.sessionCookieName)

		if err != nil {
			slog.InfoContext(r.Context(), "error getting session cookie", slog.Any("err", err))
			next.ServeHTTP(w, r)
			return
		}

		sessionID, userID, err := encoders.DecodeCookieValue(sessionCookie.Value)

		if err != nil {
			slog.InfoContext(r.Context(), "error decoding session cookie", slog.Any("err", err))
			next.ServeHTTP(w, r)
			return
		}

		slog.InfoContext(
			r.Context(),
			"Decoded session cookie",
			slog.Any("sessionID", sessionID),
			slog.Int64("userID", userID),
		)

		user, err := m.sessionStore.GetUserFromSession(r.Context(), sessionID, userID)
		if err != nil {
			slog.InfoContext(
				r.Context(),
				"error getting user from session",
				slog.Any("sessionID", sessionID),
				slog.Int64("userID", userID),
				slog.Any("err", err),
			)
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUser(ctx context.Context) *model.User {
	user := ctx.Value(UserKey)
	if user == nil {
		return nil
	}

	return user.(*model.User)
}
