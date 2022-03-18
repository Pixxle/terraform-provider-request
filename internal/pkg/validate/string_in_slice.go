package validate

import "github.com/Pixxle/terraform-provider-request/internal/pkg/utility"

func HttpMethod(val interface{}, key string) (warns []string, errs []error) {
	v := val.(string)
	if !utility.SliceContains(HTTP_VALID_METHODS, v) {
		errs = append(errs, InvalidHttpMethod{HttpMethod: v})
	}
	return
}
