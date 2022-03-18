package entity

import (
	"github.com/Pixxle/terraform-provider-request/internal/pkg/constants"
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type HTTP struct {
	url         string
	headers     map[string]string
	aws_profile string
	sigv4       bool
}

func NewHTTP(d *schema.ResourceData) (*HTTP, error) {
	url := d.Get(constants.URL).(string)
	headers := d.Get(constants.REQUEST_HEADERS).(map[string]interface{})
	/*
		query_parameters := d.Get(constants.QUERY_PARAMETERS).(map[string]interface{})
		aws_profile := d.Get(constants.AWS_PROFILE).(map[string]interface{})
		sigv4_signed := d.Get(constants.SIGV4_SIGNED).(bool)
	*/

	spew.Dump(headers)
	return &HTTP{
		url: url,
	}, nil
}
