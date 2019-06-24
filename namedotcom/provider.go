package namedotcom

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/namedotcom/go/namecom"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TOKEN", nil),
				Description: "Name.com API Token Value",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("USERNAME", nil),
				Description: "Name.com API Username",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"namedotcom_record": resourceRecord(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	username := d.Get("username").(string)
	token := d.Get("token").(string)
	nc := namecom.New(username, token)
	return nc, nil
}
