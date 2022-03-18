package main

import (
	"context"
	"flag"
	"github.com/Pixxle/terraform-provider-request/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"log"
)

func main() {

	debug := flag.Bool("debug", false, "launch provider in debug mode")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return provider.New()
		},
	}

	if *debug {
		err := plugin.Debug(context.Background(), "github.com/Pixxle/request", opts)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	plugin.Serve(opts)
}
