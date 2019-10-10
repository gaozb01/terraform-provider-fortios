package fortios

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccFortiManagerDVMInstallPolicyPackage(t *testing.T) {
	pkg_name := "dvm-pkg-" + acctest.RandString(6)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckFortiManager(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFortiManagerDVMInstallPolicyPackageConfig(pkg_name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFortiManagerDVMInstallPolicyPackageExists("fortios_fortimanager_devicemanager_install_policypackage.test3"),
					resource.TestCheckResourceAttr("fortios_fortimanager_devicemanager_install_policypackage.test3", "package_name", pkg_name),
				),
			},
		},
	})
}

func testAccCheckFortiManagerDVMInstallPolicyPackageExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found devicemanager install policypackage config file: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No devicemanager install policypackage config is set")
		}

		return nil
	}
}

func testAccFortiManagerDVMInstallPolicyPackageConfig(name string) string {
	return fmt.Sprintf(`
resource "fortios_fortimanager_firewall_security_policypackage" "test1" {
    name = "%s"
    target = "FGVM64-test"
}
resource "fortios_fortimanager_firewall_security_policy" "test2" {
    name = "policy-test"
    srcaddr = ["all"]
    srcintf = ["any"]
    dstaddr = ["all"]
    dstintf = ["any"]
    service = ["ALL"]
    action = "accept"
    schedule= ["always"]
    package_name = fortios_fortimanager_firewall_security_policypackage.test1.name
}
resource "fortios_fortimanager_devicemanager_install_policypackage" "test3" {
   package_name = fortios_fortimanager_firewall_security_policy.test2.package_name
}
`, name)
}
