package semaphoreci

import "github.com/hashicorp/terraform/helper/schema"

// dataSourceSemaphoreCIProjects returns a *schema.Resource.
func dataSourceSemaphoreCIProjects() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceSemaphoreCIProjectsRead,
		Schema: map[string]*schema.Schema{},
	}
}

// TODO
func dataSourceSemaphoreCIProjectsRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
