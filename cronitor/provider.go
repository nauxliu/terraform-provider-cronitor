package cronitor

import (
	"github.com/hashicorp/terraform/helper/schema"
)

const ApiKeyEnvName = "API_KEY"

var DefaultTimeZone string

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(ApiKeyEnvName, ""),
			},
			"default_timezone": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "UTC",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cronitor_heartbeat_monitor": resourceMonitor(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := Client{ApiKey: d.Get("api_key").(string)}
	DefaultTimeZone = d.Get("default_timezone").(string)

	return client, nil
}
