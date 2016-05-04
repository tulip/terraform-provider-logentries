package main

import (
	"github.com/donaldguy/terraform-provider-logentries/logentriesapi"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceLogset() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the log",
				Required:    true,
				ForceNew:    true,
			},
			"key": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Primary key / ID of the log",
				Computed:    true,
			},
		},
		Create: logsetCreate,
		Read:   logsetRead,
		Delete: logsetDelete,
	}
}

func logsetCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*logentriesapi.Client)
	log, err := client.CreateHost(d.Get("name").(string))

	if err != nil {
		return err
	}

	d.SetId(log.Key)
	d.Set("key", log.Key)
	return nil
}

func logsetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*logentriesapi.Client)
	log, err := client.GetHost(d.Id())

	if err != nil {
		return err
	}

	d.Set("key", log.Key)
	return nil
}

func logsetDelete(d *schema.ResourceData, meta interface{}) error {
	//client := meta.(*logentriesapi.Client)

	d.SetId("")
	return nil
}
