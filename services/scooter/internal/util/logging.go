package util

import (
	"net/http"
	"fmt"
)
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Log the request method and URL
        fmt.Printf("Request: %s %s", r.Method, r.URL.Path)

        // Call the next handler
        next.ServeHTTP(w, r)
    })
}