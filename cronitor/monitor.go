package cronitor

import (
	"github.com/hashicorp/terraform/helper/schema"
)

type Monitor struct {
	Code          string        `json:"code,omitempty"`
	Name          string        `json:"name"`
	Type          string        `json:"type"`
	Tags          []string      `json:"tags,omitempty"`
	Notifications Notifications `json:"notifications"`
	Rules         []Rule        `json:"rules"`
	Note          string        `json:"note,omitempty"`
}

type Notifications struct {
	Phones    []string `json:"phones,omitempty"`
	Webhooks  []string `json:"webhooks,omitempty"`
	Emails    []string `json:"emails,omitempty"`
	Pagerduty []string `json:"pagerduty,omitempty"`
	Slack     []string `json:"slack,omitempty"`
}

type Rule struct {
	RuleType             string `json:"rule_type"`
	Value                string `json:"value,omitempty"`
	TimeUnit             string `json:"time_unit,omitempty"`
	HoursToFollowupAlert int    `json:"hours_to_followup_alert,omitempty"`
	GraceSeconds         int    `json:"grace_seconds,omitempty"`
}

func (this Monitor) createFromResourceData(d *schema.ResourceData) (Monitor, error) {
	monitor := Monitor{
		Notifications: Notifications{},
		Rules:         []Rule{},
	}

	if attr, ok := d.GetOk("name"); ok {
		monitor.Name = attr.(string)
	}

	if attr, ok := d.GetOk("type"); ok {
		monitor.Type = attr.(string)
	}

	if attr, ok := d.GetOk("note"); ok {
		monitor.Note = attr.(string)
	}

	if attr, ok := d.GetOk("tags"); ok {
		monitor.Tags = this.convertInterfacesToStrings(attr.(*schema.Set).List())
	}

	if _, ok := d.GetOk("notifications.0"); ok {
		if attr, ok := d.GetOk("notifications.0.slack"); ok {
			monitor.Notifications.Slack = this.convertInterfacesToStrings(attr.([]interface{}))
		}

		if attr, ok := d.GetOk("notifications.0.emails"); ok {
			monitor.Notifications.Emails = this.convertInterfacesToStrings(attr.([]interface{}))
		}

		if attr, ok := d.GetOk("notifications.0.pagerduty"); ok {
			monitor.Notifications.Pagerduty = this.convertInterfacesToStrings(attr.([]interface{}))
		}

		if attr, ok := d.GetOk("notifications.0.phones"); ok {
			monitor.Notifications.Phones = this.convertInterfacesToStrings(attr.([]interface{}))
		}

		if attr, ok := d.GetOk("notifications.0.webhooks"); ok {
			monitor.Notifications.Webhooks = this.convertInterfacesToStrings(attr.([]interface{}))
		}
	}

	if attr, ok := d.GetOk("rule"); ok {
		attrRules := attr.(*schema.Set)

		for _, v := range attrRules.List() {
			rule := v.(map[string]interface{})
			monitor.Rules = append(monitor.Rules, Rule{
				RuleType:             rule["rule_type"].(string),
				Value:                rule["value"].(string),
				TimeUnit:             rule["time_unit"].(string),
				HoursToFollowupAlert: rule["hours_to_followup_alert"].(int),
				GraceSeconds:         rule["grace_seconds"].(int),
			})
		}
	}

	return monitor, nil
}

func (this Monitor) getNotificationsMapping() []interface{} {
	n := make(map[string]interface{})

	n["slack"] = this.Notifications.Slack
	n["emails"] = this.Notifications.Emails
	n["pagerduty"] = this.Notifications.Pagerduty
	n["phones"] = this.Notifications.Phones
	n["webhooks"] = this.Notifications.Webhooks

	return []interface{}{n}
}

func (this Monitor) getRulesMapping() []interface{} {
	var n []interface{}

	for _, rule := range this.Rules {
		n = append(n, map[string]interface{}{
			"rule_type":               rule.RuleType,
			"value":                   rule.Value,
			"time_unit":               rule.TimeUnit,
			"hours_to_followup_alert": rule.HoursToFollowupAlert,
			"grace_seconds":           rule.GraceSeconds,
		})
	}

	return n
}

func (this Monitor) convertInterfacesToStrings(interfaces []interface{}) []string {
	result := make([]string, len(interfaces))

	for i, v := range interfaces {
		result[i] = v.(string)
	}

	return result
}
