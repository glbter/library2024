package errors

import (
	"context"
	"log/slog"
	"net/http"
)

func ServerError(ctx context.Context, w http.ResponseWriter, err error, message string) {
	http.Error(w, message, http.StatusInternalServerError)
	slog.ErrorContext(ctx, message, slog.Any("err", err))
}
