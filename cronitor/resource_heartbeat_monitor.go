package cronitor

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceHeartbeatMonitorCreate,
		Read:   resourceHeartbeatMonitorRead,
		Update: resourceHeartbeatMonitorUpdate,
		Delete: resourceHeartbeatMonitorDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"notifications": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"emails": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"slack": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"phones": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"webhooks": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"rule": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Default:  "not_on_schedule",
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								"not_on_schedule",
								"run_ping_not_received",
								"complete_ping_not_received",
								"ran_longer_than",
								"ran_less_than",
								"run_ping_received",
								"complete_ping_received",
							}, false),
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						//  Not required for not_on_schedule rules
						"time_unit": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ValidateFunc: validation.StringInSlice([]string{
								"seconds",
								"minutes",
								"hours",
								"days",
								"weeks",
							}, false),
						},
						"hours_to_followup_alert": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"grace_seconds": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceHeartbeatMonitorCreate(d *schema.ResourceData, m interface{}) error {
	return resourceHeartbeatMonitorRead(d, m)
}

func resourceHeartbeatMonitorRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceHeartbeatMonitorUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceHeartbeatMonitorRead(d, m)
}

func resourceHeartbeatMonitorDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
