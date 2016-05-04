package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	le "github.com/logentries/le_goclient"
)

func resourceLogset() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the log",
				Required:    true,
			},
			//optional
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "nonlocation",
			},
		},
		Create: logsetCreate,
		Read:   logsetRead,
		Update: logsetUpdate,
		Delete: logsetDelete,
	}
}

func setLogSetVals(d *schema.ResourceData, logSet *le.LogSet) {
	d.Set("name", logSet.Name)
	d.Set("location", logSet.Location)
}

func logsetCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*le.Client)
	logSet, err := client.LogSet.Create(le.LogSetCreateRequest{
		Name:     d.Get("name").(string),
		Location: d.Get("location").(string),
	})

	if err != nil {
		return err
	}
	d.SetId(logSet.Key)
	setLogSetVals(d, logSet)
	return nil
}

func logsetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*le.Client)
	logSet, err := client.LogSet.Read(le.LogSetReadRequest{Key: d.Id()})
	if err != nil {
		return err
	}

	setLogSetVals(d, logSet)
	return nil
}

func logsetUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*le.Client)
	logSet, err := client.LogSet.Update(le.LogSetUpdateRequest{
		Key:      d.Id(),
		Name:     d.Get("name").(string),
		Location: d.Get("location").(string),
	})

	if err != nil {
		return err
	}

	setLogSetVals(d, logSet)
	return nil
}

func logsetDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*le.Client)
	err := client.LogSet.Delete(le.LogSetDeleteRequest{Key: d.Id()})

	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
