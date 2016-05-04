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
			"in_logset_with_key": &schema.Schema{
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
			"retention": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "-1",
			},
			//strictly Optional
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			//strictly computed
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
			"follow": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Create: logCreate,
		Read:   logRead,
		Update: logUpdate,
		Delete: logDelete,
	}
}

func setLogVals(d *schema.ResourceData, log *le.Log) {
	d.Set("name", log.Name)

	d.Set("key", log.Key)
	d.Set("token", log.Token)

	d.Set("type", log.Type)
	d.Set("retention", log.Retention)

	d.Set("filename", log.Filename)
	d.Set("source", log.Source)
	d.Set("follow", log.Follow)

}

func logCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*le.Client)

	req := le.LogCreateRequest{}
	req.Name = d.Get("name").(string)
	req.LogSetKey = d.Get("in_logset_with_key").(string)
	req.Retention = d.Get("retention").(string)
	req.Source = d.Get("source").(string)

	if v, ok := d.GetOk("filename"); ok {
		req.Filename = v.(string)
	}
	if v, ok := d.GetOk("source"); ok {
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
		Retention: d.Get("retention").(string),
	})
	if err != nil {
		return err
	}

	setLogVals(d, log)
	return nil
}

func logDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*le.Client)
	req := le.LogDeleteRequest{}
	req.Key = d.Id()
	req.LogSetKey = d.Get("in_logset_with_key").(string)

	err := client.Log.Delete(req)

	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
