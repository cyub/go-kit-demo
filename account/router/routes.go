package router

import (
	"net/http"

	"github.com/cyub/go-kit-demo/account/service"
)

// Route define route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is array of route
type Routes []Route

var routes = Routes{
	Route{
		Name:        "Register",
		Method:      "POST",
		Pattern:     "/register",
		HandlerFunc: service.Register,
	},
	Route{
		Name:        "Show",
		Method:      "Get",
		Pattern:     "/show/{id:[0-9]+}",
		HandlerFunc: service.Show,
	},
}
