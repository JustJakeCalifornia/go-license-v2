package routes

import (
	"net/http"

	"github.com/trite8q1/go-license-v2/keygen-service/core/routes/rest"
)

const (
	APIVersionV1 = "v1"
	APIVersionV2 = "v2"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	APIVersion  string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		APIVersionV1,
		rest.Index,
	},
	Route{
		"Health",
		"GET",
		"/health",
		APIVersionV1,
		rest.Health,
	},
	Route{
		"GenerateLicense",
		"GET",
		"/licenses",
		APIVersionV1,
		rest.GetLicenses,
	},
}
