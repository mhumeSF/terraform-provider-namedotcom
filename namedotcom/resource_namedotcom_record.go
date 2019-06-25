package namedotcom

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/namedotcom/go/namecom"
)

func resourceRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceRecordCreate,
		Read:   resourceRecordRead,
		Update: resourceRecordUpdate,
		Delete: resourceRecordDelete,

		Schema: map[string]*schema.Schema{
			"domain_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "DomainName is the zone that the record belongs to",
			},
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Host is the hostname relative to the zone",
			},
			"record_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Type is one of the following: A, AAAA, ANAME, CNAME, MX, NS, SRV, or TXT",
			},
			"answer": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Answer is either the IP address for A or AAAA records",
			},
		},
	}
}

func resourceRecordCreate(d *schema.ResourceData, m interface{}) error {
	// createRecord
	client := m.(*NameDotComClient.conn)
	record := &namecom.Record{
		DomainName: d.Get("domain_name").(string),
		Host:       d.Get("host").(string),
		Type:       d.Get("record_type").(string),
		Answer:     d.Get("answer").(string),
	}
	client.CreateRecord(record)
	return nil
}

func resourceRecordRead(d *schema.ResourceData, m interface{}) error {
	// readRecord
	return nil
}

func resourceRecordUpdate(d *schema.ResourceData, m interface{}) error {
	// updateRecord
	return nil
}

func resourceRecordDelete(d *schema.ResourceData, m interface{}) error {
	// deleteRecord
	return nil
}
