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
			State: schema.ImportStatePassthrough,
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
						"templates": &schema.Schema{
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
								"not_run_in",
								"run_ping_not_received",
								"not_completed_in",
								"complete_ping_not_received",
								"ran_longer_than",
								"completed_in_under",
								"ran_less_than",
								"has_started",
								"run_ping_received",
								"has_completed",
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
				Default:  "",
			},
			"timezone": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				DefaultFunc: func() (interface{}, error) {
					return DefaultTimeZone, nil
				},
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

	code, err := client.Create(monitor)

	if err != nil {
		return err
	}

	d.SetId(*code)

	return resourceHeartbeatMonitorRead(d, m)
}

func resourceHeartbeatMonitorRead(d *schema.ResourceData, m interface{}) error {
	client := m.(Client)

	monitor, err := client.Read(d.Id())

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

	if err = d.Set("timezone", monitor.Timezone); err != nil {
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

	if err := client.Update(d.Id(), monitor); err != nil {
		return err
	}

	return resourceHeartbeatMonitorRead(d, m)
}

func resourceHeartbeatMonitorDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(Client)

	return client.Delete(d.Id())
}
