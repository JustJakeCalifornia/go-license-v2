package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/trite8q1/go-license-v2/keygen-service/internal/middleware"
	"github.com/trite8q1/go-license-v2/keygen-service/utils"
)

// NewRouter
func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	middleware.InitializeMiddleware(router)
	middleware.InitializeCors(router)

	router.Route(APIVersionV1, func(r chi.Router) {
		getRoutes(router)
	})

	return router
}

func getRoutes(router *chi.Mux) []Route {
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = utils.Logger(handler, route.Name)

		router.Method(route.Method, route.APIVersion+route.Pattern, handler)
		fmt.Println("hier drinnen")
	}
	return routes
}
