package middleware

import (
	"context"
	"net/http"

	"database/sql"
)

func dbMiddleware(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Inject the db parameter into the request context
			ctx := context.WithValue(r.Context(), "db", db)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

