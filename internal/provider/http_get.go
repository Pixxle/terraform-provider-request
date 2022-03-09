package provider

import (
	"context"
	"github.com/Pixxle/terraform-provider-request/internal/pkg/constants"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func httpGet() *schema.Resource {
	return &schema.Resource{
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

func dataSourceHTTPRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostic) {
	url := d.Get(constants.URL).(string)
	headers := d.Get(constants.REQUEST_HEADERS).(map[string]interface{})
	query_parameters := d.Get(constants.QUERY_PARAMETERS).(map[string]interface{})
	aws_profile := d.Get(constants.AWS_PROFILE).(map[string]interface{})
	sigv4_signed := d.Get(constants.SIGV4_SIGNED).(bool)
}
