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
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				if err := resourceHeartbeatMonitorRead(d, m); err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
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
						"pagerduty": &schema.Schema{
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
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_type": &schema.Schema{
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
							Optional: true,
						},
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
							Default:  8,
						},
						"grace_seconds": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
					},
				},
			},
			"tags": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceHeartbeatMonitorCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(Client)

	monitor, err := Monitor{}.createFromResourceData(d)

	if err != nil {
		return err
	}

	monitor.Type = "heartbeat"

	code, err := client.create(monitor)

	if err != nil {
		return err
	}

	d.SetId(*code)

	return resourceHeartbeatMonitorRead(d, m)
}

func resourceHeartbeatMonitorRead(d *schema.ResourceData, m interface{}) error {
	client := m.(Client)

	monitor, err := client.read(d.Id())

	if err != nil {
		return err
	}

	if err = d.Set("name", monitor.Name); err != nil {
		return err
	}

	if err = d.Set("tags", monitor.Tags); err != nil {
		return err
	}

	if err = d.Set("note", monitor.Note); err != nil {
		return err
	}

	if err = d.Set("notifications", monitor.getNotificationsMapping()); err != nil {
		return err
	}

	if err = d.Set("rule", monitor.getRulesMapping()); err != nil {
		return err
	}

	return nil
}

func resourceHeartbeatMonitorUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(Client)

	monitor, err := Monitor{}.createFromResourceData(d)

	if err != nil {
		return err
	}

	monitor.Type = "heartbeat"

	if err := client.update(d.Id(), monitor); err != nil {
		return err
	}

	return resourceHeartbeatMonitorRead(d, m)
}

func resourceHeartbeatMonitorDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(Client)

	return client.delete(d.Id())
}
