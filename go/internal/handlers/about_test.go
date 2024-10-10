package handlers

import (
	"bytes"
	"context"
	"library/internal/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAboutHandler(t *testing.T) {

	testCases := []struct {
		name               string
		expectedStatusCode int
		expectedBody       []byte
	}{
		{
			name:               "render successfully",
			expectedStatusCode: http.StatusOK,
			expectedBody:       []byte("Library - About"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			assert := assert.New(t)

			handler := NewAboutHandler()

			req, err := http.NewRequest("GET", "/about", nil)
			assert.NoError(err)

			value := middleware.Nonces{
				Htmx:            "nonce-1234",
				ResponseTargets: "nonce-5678",
				Tw:              "nonce-9101",
				HtmxCSSHashes: []string{
					"sha256-VV50kYRP+CuHcIPnOJ/AV+Q2C0IGVX7AczE6/dxv078=",
					"sha256-cNCcTUjHx0E9D/2VhCEWby6Bd/Ow9sHRp65Rf9s3J68=",
					"sha256-9Ccj/TQB4XGZUjlcvis7DszVKQz1ppCoyRybka1oFIA=",
					"sha256-bsV5JivYxvGywDAZ22EZJKBFip65Ng9xoJVLbBg7bdo=",
				},
			}
			ctx := context.WithValue(req.Context(), middleware.NonceKey, value)
			req = req.WithContext(ctx)

			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			assert.Equal(tc.expectedStatusCode, rr.Code, "handler returned wrong status code: got %v want %v", rr.Code, tc.expectedStatusCode)

			assert.True(bytes.Contains(rr.Body.Bytes(), tc.expectedBody), "handler returned unexpected body: got %v want %v", rr.Body.String(), tc.expectedBody)

		})

	}

}
