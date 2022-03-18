package provider

import (
	"context"
	"fmt"
	"github.com/Pixxle/terraform-provider-request/internal/pkg/constants"
	"github.com/Pixxle/terraform-provider-request/internal/pkg/entity"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func httpGet() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceHTTPRead,
		Schema: map[string]*schema.Schema{
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
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			constants.SIGV4_SIGNED: {
				Type:     schema.TypeBool,
				Optional: true,
			},
		}}
}

func dataSourceHTTPRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	_, err := entity.NewHTTP(d)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Generate HTTP Request failed due to %s", err.Error()),
		})
		return
	}
	return
}
