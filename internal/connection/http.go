package request

import (
	"fmt"
	"github.com/Pixxle/terraform-provider-request/internal/constants"
	"github.com/Pixxle/terraform-provider-request/internal/utility"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io"
	"net/http"
	"strings"
)

func NewHTTP(d *schema.ResourceData) (*http.Request, error) {

	httpMethod := d.Get(constants.HTTP_METHOD).(string)
	url := d.Get(constants.URL).(string)
	bodyContent := d.Get(constants.BODY_CONTENT).(string)

	var bcR io.Reader
	if len(bodyContent) != 0 {
		bcR = strings.NewReader(bodyContent)
	}

	request, err := http.NewRequest(httpMethod, url, bcR)
	if err != nil {
		return nil, fmt.Errorf("failed to generate http request %v", err)
	}

	headers := d.Get(constants.REQUEST_HEADERS).(map[string]interface{})
	for k, v := range headers {
		request.Header.Add(k, v.(string))
	}

	queryParameters := d.Get(constants.QUERY_PARAMETERS).(map[string]interface{})
	value := request.URL.Query()
	for k, v := range queryParameters {
		value.Add(k, v.(string))
	}
	request.URL.RawQuery = value.Encode()

	sigv4Signed := d.Get(constants.SIGV4_SIGNED).(bool)
	sigv4Service := d.Get(constants.SIGV4_SERVICE).(string)
	awsProfile := d.Get(constants.AWS_PROFILE).(string)
	awsRegion := d.Get(constants.AWS_REGION).(string)

	if sigv4Signed {
		if err = utility.SigV4SignRequest(request, awsProfile, awsRegion, sigv4Service); err != nil {
			return nil, err
		}
	}

	return request, nil
}
