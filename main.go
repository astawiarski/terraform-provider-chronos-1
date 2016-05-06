package main

import (
	"terraform-provider-chronos/chronos"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: chronos.Provider,
	})
}
