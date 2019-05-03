# Terraform Provider for Cronitor

## Provider configuration

| Property | Description                 |  Type  | Required |
|----------|-----------------------------|--------|----------|
| api_key  | The API Key                 | string | true     |


## Example

```HCL
provider "cronitor" {
  api_key = "{your api key}"
}

resource "cronitor_heartbeat_monitor" "example" {
    name = "foobar"

    notifications = {
        webhooks = ["https://webhook.url"]
        slack = ["https://slack.incoming.webhook.url"]
    }

    rule = {
        value = "* * * * * *"
        grace_seconds = 30
    }
    
    rule = {
        rule_type = "run_ping_not_received"
    }
    
    rule = {
        rule_type = "ran_less_than"
        time_unit = "minutes"
        value = 10
    }

    tags = ["foo", "bar"]

    note = "heartbeat monitor"
}
```

## Import

Monitor can be imported using the code, e.g.

```
$ terraform import cronitor_heartbeat_monitor.example EaB9B2
```
