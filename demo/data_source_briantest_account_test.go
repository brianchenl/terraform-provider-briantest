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
					resource.TestCheckResourceAttrSet("data.briantest_account.current", "id"),
					resource.TestCheckResourceAttrSet("data.briantest_account.current", "name"),
				),
			},
		},
	})
}

const testAccCheckBrianAccountDataSourceBasic = `
data "briantest_account" "current" {
	name = "ecs"
}
`
