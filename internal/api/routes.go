package api

import "net/http"

const (
	routeEntry = "/api/entry"
	routeTable = "/api/table"
)

func MountRoutes(mux *http.ServeMux) *http.ServeMux {
	return mux
}
