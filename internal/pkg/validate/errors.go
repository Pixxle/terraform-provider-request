package validate

import "fmt"

type InvalidHttpMethod struct {
	HttpMethod string
}

func (e InvalidHttpMethod) Error() string {
	return fmt.Sprintf("invalid http method %s provided, valid options are %+v", e.HttpMethod, HTTP_VALID_METHODS)
}
