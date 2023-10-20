package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	ctx "github.com/tehrelt/url-shortener/internal/app/api/context"
)

func SetRequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()

		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctx.CtxKeyRequestID, id)))
	})
}
