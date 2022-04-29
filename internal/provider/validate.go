package provider

import (
	"github.com/Pixxle/terraform-provider-request/internal/utility"
	"net/http"
)

var (
	HTTP_VALID_METHODS = []string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodHead, http.MethodOptions, http.MethodPatch, http.MethodTrace, http.MethodConnect}
)

func ValidateHTTPMethod(val interface{}, key string) (warns []string, errs []error) {
	v := val.(string)
	if !utility.SliceContains(HTTP_VALID_METHODS, v) {
		errs = append(errs, InvalidHttpMethod{HttpMethod: v})
	}
	return
}
