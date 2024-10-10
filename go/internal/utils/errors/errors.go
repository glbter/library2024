package errors

import (
	"context"
	"log/slog"
	"net/http"
)

func ServerError(ctx context.Context, w http.ResponseWriter, err error, message string, args ...any) {
	http.Error(w, message, http.StatusInternalServerError)
	slog.ErrorContext(ctx, message, append([]any{slog.Any("err", err)}, args...)...)
}
