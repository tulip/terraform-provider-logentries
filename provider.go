package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	le "github.com/logentries/le_goclient"
)

func provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"account_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LOGENTRIES_KEY", nil),
				Description: "Logentries account key ",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"logentries_logset": resourceLogset(),
			"logentries_log":    resourceLog(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return le.NewClient(d.Get("account_key").(string)), nil
}
