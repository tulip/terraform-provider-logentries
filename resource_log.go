package main

import (
	"github.com/donaldguy/terraform-provider-logentries/logentriesapi"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceLog() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the log",
				Required:    true,
				ForceNew:    true,
			},
			"in_logset_with_key": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Key of logset under which to parent the log",
				Required:    true,
				ForceNew:    true,
			},
			"key": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Primary key / ID of the log",
				Computed:    true,
			},
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Token used to submit data to log",
				Computed:    true,
			},
		},
		Create: logCreate,
		Read:   logRead,
		Delete: logDelete,
	}
}

func logCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*logentriesapi.Client)
	log, err := client.CreateLog(
		d.Get("in_logset_with_key").(string),
		d.Get("name").(string),
	)

	if err != nil {
		return err
	}

	d.SetId(log.Key)
	d.Set("key", log.Key)
	d.Set("token", log.Token)
	return nil
}

func logRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*logentriesapi.Client)
	log, err := client.GetLog(d.Id())

	if err != nil {
		return err
	}

	d.Set("name", log.Name)
	d.Set("key", log.Key)
	d.Set("token", log.Token)
	return nil
}

func logDelete(d *schema.ResourceData, meta interface{}) error {
	//client := meta.(*logentriesapi.Client)

	d.SetId("")
	return nil
}
