package middlewares

import (
	"net/http"
	"strings"
)

func CorsMiddleware(routes map[string]map[string]http.HandlerFunc) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			allowedMethods := []string{}

			if methodsMap, ok := routes[r.URL.Path]; ok {
				for method := range methodsMap {
					allowedMethods = append(allowedMethods, method)
				}
			}

			// Always allow OPTIONS for preflight
			allowedMethods = append(allowedMethods, http.MethodOptions)

			// Set CORS headers
			w.Header().Set("Access-Control-Allow-Origin", "*") // or your domain
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, X-CSRF-Token")
			w.Header().Set("Access-Control-Expose-Headers", "X-CSRF-Token")
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(allowedMethods, ", "))

			if r.Method == http.MethodOptions {
				// Respond immediately to OPTIONS requests
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
