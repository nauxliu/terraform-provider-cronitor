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
    }

    rule = {
        value = "* * * * * *"
    }

    tags = ["fo", "bar"]

    note = "heartbeat monitor"
}
```

## Import

Monitor can be imported using the code, e.g.

```
$ terraform import cronitor_heartbeat_monitor.example EaB9B2
```
