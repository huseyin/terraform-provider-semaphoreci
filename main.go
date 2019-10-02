package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/huseyin/terraform-provider-semaphoreci/semaphoreci"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: semaphoreci.Provider})
}
