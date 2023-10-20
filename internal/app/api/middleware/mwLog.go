package middleware

import (
	"log/slog"
	"net/http"
	"time"

	ctx "github.com/tehrelt/url-shortener/internal/app/api/context"
)

func LogMiddleware(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		// log = log.With(
		// 	slog.String("component", "middleware/mwLog"),
		// )

		// log.Info("log middleware enabled")

		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := log.With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("request-id", r.Context().Value(ctx.CtxKeyRequestID).(string)),
			)

			rw := &ResponseWriter{w, http.StatusOK}

			t := time.Now()

			defer func() {
				entry.Info(
					"request completed",
					slog.Int("status", rw.code),
					slog.String("duration", time.Since(t).String()),
				)
			}()

			next.ServeHTTP(rw, r)
		}

		return http.HandlerFunc(fn)
	}
}
