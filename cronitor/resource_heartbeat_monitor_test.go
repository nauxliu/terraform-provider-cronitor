package cronitor

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestHeartbeatMonitor_basic(t *testing.T) {
	name := "Test_" + acctest.RandStringFromCharSet(7, acctest.CharSetAlphaNum)
	rn := fmt.Sprintf("cronitor_heartbeat_monitor.%s", name)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testHeartbeatMonitor_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "note", ""),
					resource.TestCheckResourceAttrSet(rn, "timezone"),
				),
			},
			{
				Config: testHeartbeatMonitor_complete(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(rn, "name", name),
					resource.TestCheckResourceAttr(rn, "note", "Hello"),
					resource.TestCheckResourceAttr(rn, "timezone", "Asia/Chongqing"),
				),
			},
		},
	})
}

func testHeartbeatMonitor_basic(name string) string {
	return fmt.Sprintf(`
resource "cronitor_heartbeat_monitor" "%s" {
  name = "%s"

  notifications {
	webhooks = ["https://webhook.url"]
  }

  rule {
    value = "* * * * * *"
  }
}`, name, name)
}

func testHeartbeatMonitor_complete(name string) string {
	return fmt.Sprintf(`
resource "cronitor_heartbeat_monitor" "%s" {
  name = "%s"

  notifications {
	webhooks = ["https://webhook.url"]
  }

  rule {
    value = "* * * * * *"
  }

  note = "Hello"

  timezone = "Asia/Chongqing"
}`, name, name)
}
