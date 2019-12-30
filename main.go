package main

import (
	"github.com/terraform-providers/terraform-provider-onepassword/onepassword"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: onepassword.Provider,
	})
}
