package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/trite8q1/go-license-v2/keygen-service/utils"
)

// NewRouter
func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		for _, route := range routes {
			var handler http.Handler

			handler = route.HandlerFunc
			handler = utils.Logger(handler, route.Name)

			router.Method(route.Method, route.Pattern, handler)
		}
	})

	return router
}

// func getRoutes(apiVersion string, router *chi.Router) []Route {
// 	for _, route := range routes {
// 		var handler http.Handler

// 		handler = route.HandlerFunc
// 		handler = utils.Logger(handler, route.Name)

// 		if route.APIVersion == apiVersion {
// 			router.r

// 		}
// 	}
// 	return routes
// }
