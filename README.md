# Terraform Provider for Cronitor

[![Build Status](https://github.com/nauxliu/terraform-provider-cronitor/workflows/Lint/badge.svg)](https://github.com/nauxliu/terraform-provider-cronitor/actions) 
[![LICENSE](https://img.shields.io/github/license/nauxliu/terraform-provider-cronitor)](https://github.com/nauxliu/terraform-provider-cronitor/blob/master/LICENSE) 

Allows you to manage cronitor monitors.

## How to use this provider

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```
terraform {
  required_providers {
    cronitor = {
      source = "nauxliu/cronitor"
      version = ">=1.0.3"
    }
  }
}
```

## Provider configuration

The following arguments are supported in the provider block:

`api_key` - (Required) API key of Cronitor  
`default_timezone` - (Optional) Default timezone for all monitors

# Resources

## cronitor_heartbeat_monitor


### Example Usage

```HCL
provider "cronitor" {
  api_key = "{your api key}"
}

resource "cronitor_heartbeat_monitor" "example" {
    name = "foobar"

    notifications {
        webhooks = ["https://webhook.url"]
        slack = ["https://slack.incoming.webhook.url"]
    }

    rule {
        value = "* * * * * *"
        grace_seconds = 30
    }
    
    rule {
        rule_type = "run_ping_not_received"
    }
    
    rule {
        rule_type = "ran_less_than"
        time_unit = "minutes"
        value = 10
    }

    tags = ["foo", "bar"]

    note = "heartbeat monitor"
}
```

### Argument Reference

The following arguments are supported:

* `name` - (Required) the name of your monitor. Note: All of your monitors must have a unique name.
* `notifications` - (Required) where/how you wish to be contacted when a monitor's alerting is triggered. The following key/value pairs are all options, at least one of which must not be empty. Note: When extending notification template(s), passing an empty array will overload the templated notification settings for that key. 
  * `templates` (Optional) - ordered list of templates (by key) to extend from. Merged from left-to-right, and then with any monitor-specific template settings defined here.
  * `emails` (Optional) - list of emails to send alerts to
  * `slack` (Array) - list of slack webhook URLs (found on account settings page)
  * `pagerduty` (Optional) - list of pagerduty keys (found on account settings page)
  * `phones` (Optional) - list of phone numbers to send SMS alerts to
  * `webhooks` (Optional) - list of URLs (prefixed with http:// or https://) to callback to 
* `rule` (Required) - When creating and updating a monitor you must specify the rules that will trigger alerts to be sent.
  * `rule_type` (Optional) - Options are:
    * `not_on_schedule` // default
    * `run_ping_not_received`
    * `complete_ping_not_received`
    * `ran_longer_than`
    * `ran_less_than`
    * `run_ping_received`
    * `complete_ping_received`
  * `value` (Optional) - For `not_on_schedule` rules, this should be a cron expression like `*/10 * * * 1-5`. For `run_ping_received` and `complete_ping_received`, no value or time_unit is accepted. For all other rule types, this should be a number that is combined with with `time_unit` to specify a time interval.
  * `time_unit` (Optional) - Not required for not_on_schedule rules. Options are:
    * `seconds`
    * `minutes`
    * `hours`
    * `days`
    * `weeks`
  * `hours_to_followup_alert` (Optional) - how long to wait between sending you follow up alerts. By default Cronitor will wait 8 hours before sending a second round of alerts. The minimum value that you may set this to is 1 hour.
  * `grace_seconds` (Optional)  - Specify a grace period for evaluation of this rule. For not_on_schedule rules, this is used when evaluating start time and total runtime duration.
* `tags` (Optional) - A list of tags. Each tag must be a string <= 50 chars.
* `note` (Optional) - a note that you would like to have included in alerts. It's useful if you need to include context/tips for anyone receiving these alerts. Note that due to the size limits on SMS messages it will not be included in those alerts.
* `timezone` (Optional) - Override the provider's default timezone

### Import

Monitor can be imported using the code, e.g.

```
$ terraform import cronitor_heartbeat_monitor.example EaB9B2
```
