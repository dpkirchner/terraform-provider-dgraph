package main

import (
	"github.com/dgraph-io/dgo/v2"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"livingit.de/code/tf-dgraph/resources"
	"livingit.de/code/tf-dgraph/resources/predicate"
)

var client *dgo.Dgraph

func resourcePredicate() *schema.Resource {
	return &schema.Resource{
		Create: resources.Retry(predicate.Create),
		Read:   predicate.Read,
		Update: resources.Retry(predicate.Update),
		Delete: resources.Retry(predicate.Delete),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"array": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
			},
			"lang": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
			},
			"index": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
			},
			"tokenizer": {
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
				Description: "Required when index is true",
			},
			"edge_count": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
			},
			"reverse": {
				Type:     schema.TypeBool,
				Required: false,
				Optional: true,
			},
		},
	}
}
