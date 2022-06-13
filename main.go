package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	demo "terraform-provider-briantest/demo"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: demo.Provider,
	})
}
