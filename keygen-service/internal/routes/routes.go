package routes

import (
	"net/http"

	"github.com/trite8q1/go-license-v2/keygen-service/internal/routes/rest"
)

const (
	APIPrefix    = "/api"
	APIVersionV1 = APIPrefix + "/v1"
	APIVersionV2 = APIPrefix + "/v2"
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
		"Health",
		"GET",
		"/health",
		APIVersionV1,
		rest.Health,
	},
	Route{
		"GenerateLicenseV1",
		"GET",
		"/licenses",
		APIVersionV1,
		rest.GetLicenses,
	},
	Route{
		"GenerateLicenseV2",
		"GET",
		"/licenses",
		APIVersionV2,
		rest.GetLicensesV2,
	},
}
