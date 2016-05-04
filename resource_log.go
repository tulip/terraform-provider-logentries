package main

import (
	le "github.com/logentries/le_goclient"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceLog() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			//required
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the log",
				Required:    true,
			},
			"logset_id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Key of logset under which to parent the log",
				Required:    true,
				ForceNew:    true,
			},
			//with defaults
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "token",
			},
			"retention_period": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "-1",
			},
			//strictly optional
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			//strictly computed
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Token used to submit data to log",
				Computed:    true,
			},
		},
		Create: logCreate,
		Read:   logRead,
		Update: logUpdate,
		Delete: logDelete,
	}
}

func setLogVals(d *schema.ResourceData, log *le.Log) {
	d.Set("token", log.Token)

	d.Set("name", log.Name)
	d.Set("source", log.Source)
	d.Set("retention_period", log.Retention)
	d.Set("type", log.Type)
}

func logCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*le.Client)

	req := le.LogCreateRequest{}
	req.Name = d.Get("name").(string)
	req.LogSetKey = d.Get("logset_id").(string)
	req.Retention = d.Get("retention_period").(string)
	req.Source = d.Get("source").(string)

	if v, ok := d.GetOk("type"); ok {
		req.Source = v.(string)
	}

	log, err := client.Log.Create(req)

	if err != nil {
		return err
	}

	d.SetId(log.Key)
	setLogVals(d, log)
	return nil
}

func logRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*le.Client)

	log, err := client.Log.Read(le.LogReadRequest{Key: d.Id()})
	if err != nil {
		return err
	}

	setLogVals(d, log)
	return nil
}

func logUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*le.Client)

	log, err := client.Log.Update(le.LogUpdateRequest{
		Key:       d.Id(),
		Name:      d.Get("name").(string),
		Type:      d.Get("type").(string),
		Source:    d.Get("source").(string),
		Retention: d.Get("retention_period").(string),
	})
	if err != nil {
		return err
	}

	setLogVals(d, log)
	return nil
}

func logDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*le.Client)

	err := client.Log.Delete(le.LogDeleteRequest{
		Key:       d.Id(),
		LogSetKey: d.Get("logset_id").(string),
	})

	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
