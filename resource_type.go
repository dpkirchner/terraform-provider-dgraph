package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"livingit.de/code/tf-dgraph/resources"
	"livingit.de/code/tf-dgraph/resources/dtype"
)

func resourceType() *schema.Resource {
	return &schema.Resource{
		Create: resources.Retry(dtype.Create),
		Read:   dtype.Read,
		Update: resources.Retry(dtype.Update),
		Delete: resources.Retry(dtype.Delete),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fields": {
				Type:     schema.TypeMap,
				Required: true,
			},
		},
	}
}
