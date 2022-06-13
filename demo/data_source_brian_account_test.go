package demo


import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)


func TestAccBrianAccountDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckBrianAccountDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.brian_account.current", "id"),
					resource.TestCheckResourceAttrSet("data.brian_account.current", "name"),
				),
			},
		},
	})
}

const testAccCheckBrianAccountDataSourceBasic = `
data "brian_account" "current" {
	name = "ecs"
}
`
