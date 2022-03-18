package validate

import "net/http"

var (
	HTTP_VALID_METHODS = []string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodHead, http.MethodOptions, http.MethodPatch, http.MethodTrace, http.MethodConnect}
)
