package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/sedicii/terraform-provider-debug/debug"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: debug.Provider})
}
