package provider

import (
	"context"
	"github.com/Pixxle/terraform-provider-request/internal/pkg/constants"
	"github.com/Pixxle/terraform-provider-request/internal/pkg/request"
	"github.com/Pixxle/terraform-provider-request/internal/pkg/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
)

func httpGet() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHTTPRequest,
		Schema: map[string]*schema.Schema{
			constants.HTTP_METHOD: {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      http.MethodGet,
				ValidateFunc: validate.HttpMethod,
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
		}}
}

func dataSourceHTTPRequest(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	request, err := request.NewHTTP(d)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to generate HTTP(S) Request",
			Detail:   err.Error(),
		})
		return
	}

	client := http.Client{}
	_, err = client.Do(request)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error while sending HTTP request",
			Detail:   err.Error(),
		})
	}

	return
}
