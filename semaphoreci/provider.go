package semaphoreci

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/semaphoreci/cli/config"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_host": {
				Type:        schema.TypeString,
				Optional:    false,
				DefaultFunc: schema.EnvDefaultFunc("SEMAPHORECI_API_HOST", nil),
				Description: "The API Host for operations.",
			},
			"api_token": {
				Type:        schema.TypeString,
				Optional:    false,
				DefaultFunc: schema.EnvDefaultFunc("SEMAPHORECI_API_TOKEN", nil),
				Description: "The API Token for operations.",
			},
		},
	}
}

// providerConfigure returns a configured API client.
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config.SetHost(d.Get("api_host").(string))
	config.SetAuth(d.Get("api_token").(string))

	return nil, nil
}
