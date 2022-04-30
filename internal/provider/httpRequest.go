package provider

import (
	"context"
	"encoding/json"
	"github.com/Pixxle/terraform-provider-request/internal/connection"
	"github.com/Pixxle/terraform-provider-request/internal/constants"
	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io/ioutil"
	"net/http"
	"strings"
)

func httpRequest() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHTTPRequest,
		Schema: map[string]*schema.Schema{
			constants.HTTP_METHOD: {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      http.MethodGet,
				ValidateFunc: ValidateHTTPMethod,
			},
			constants.URL: {
				Type:     schema.TypeString,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			constants.REQUEST_HEADERS: {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			constants.QUERY_PARAMETERS: {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			constants.AWS_PROFILE: {
				Type:     schema.TypeString,
				Optional: true,
			},
			constants.AWS_REGION: {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{constants.SIGV4_SIGNED},
			},
			constants.SIGV4_SIGNED: {
				Type:     schema.TypeBool,
				Optional: true,
			},
			constants.SIGV4_SERVICE: {
				Type:         schema.TypeString,
				Optional:     true,
				RequiredWith: []string{constants.SIGV4_SIGNED},
			},
			constants.BODY_CONTENT: {
				Type:     schema.TypeString,
				Optional: true,
			},
			constants.BODY: {
				Type:     schema.TypeMap,
				Computed: true,
			},
			constants.RESPONSE_CODE: {
				Type:     schema.TypeInt,
				Computed: true,
			},
		}}
}

func dataSourceHTTPRequest(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	request, err := connection.NewHTTP(d)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to generate HTTP(S) Request",
			Detail:   err.Error(),
		})
		return
	}

	client := http.Client{}
	res, err := client.Do(request)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error while sending HTTP(S) Request",
			Detail:   err.Error(),
		})
		return diags
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error while reading HTTP(S) response body",
			Detail:   err.Error(),
		})
		return diags
	}
	defer res.Body.Close()
	i, _ := uuid.GenerateUUID()
	d.SetId(i)

	d.Set(constants.RESPONSE_CODE, res.StatusCode)
	if strings.Contains(res.Header.Get(constants.CONTENT_TYPE), constants.APPLICATION_JSON) {
		var i interface{}
		err = json.Unmarshal(body, &i)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Error while unmarshalling JSON response body",
				Detail:   err.Error(),
			})
			return diags
		}
		d.Set(constants.BODY, i)
	} else {
		d.Set(constants.BODY, &struct {
			Value string
		}{
			Value: string(body),
		})
	}

	return
}
