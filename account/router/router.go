package router

import (
	"github.com/gorilla/mux"
)

// NewRouter create a pointer to mux.Router
func New() *mux.Router {
	r := mux.NewRouter()
	for _, route := range routes {
		r.Path(route.Pattern).
			Methods(route.Method).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return r
}
