package semaphoreci

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/semaphoreci/cli/api/client"
)

// dataSourceSemaphoreCIProjects returns a *schema.Resource.
func dataSourceSemaphoreCIProjects() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSemaphoreCIProjectsRead,

		Schema: map[string]*schema.Schema{
			"projects": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

// TODO
func dataSourceSemaphoreCIProjectsRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Reading projects")

	projectAPI := client.NewProjectV1AlphaApi()
	_, err := projectAPI.ListProjects()
	if err != nil {
		return fmt.Errorf("error reading projects from SemaphoreCI: %s", err)
	}

	return nil
}
