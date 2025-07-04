package app

import (
	"net/http"
	middleware "placeholder/dev/features/middlewares"
	"strings"
)

type Router struct {
	Mux               *http.ServeMux
	Routes            map[string]map[string]http.HandlerFunc
	GlobalMiddlewares []middleware.Middleware
}

func CreateRouter() Router {
	return Router{
		Mux:    http.NewServeMux(),
		Routes: make(map[string]map[string]http.HandlerFunc),
	}
}

// This function register new routes into the router
func (router *Router) RegisterRoutes(routePATH string, method string, handler http.HandlerFunc) {
	if router.Routes[routePATH] == nil {
		router.Routes[routePATH] = make(map[string]http.HandlerFunc)
		router.Mux.HandleFunc(routePATH, func(w http.ResponseWriter, r *http.Request) {
			methodsMap := router.Routes[routePATH]

			if h, ok := methodsMap[r.Method]; ok {
				h.ServeHTTP(w, r)
				return
			}

			allowedMethods := make([]string, 0, len(methodsMap))
			for m := range methodsMap {
				allowedMethods = append(allowedMethods, m)
			}

			w.Header().Set("Allow", strings.Join(allowedMethods, ", "))
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		})
	}

	router.Routes[routePATH][method] = handler
}

// This function add new globalmiddlewares to the chain
func (router *Router) Use(middlewares ...middleware.Middleware) {
	router.GlobalMiddlewares = append(router.GlobalMiddlewares, middlewares...)
}

// This function will handle the incoming request based on the globalmiddlewares
func (router *Router) Handle() http.Handler {
	return middleware.Chain(router.Mux, router.GlobalMiddlewares...)
}
