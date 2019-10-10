---
layout: "fortios"
page_title: "FortiOS: fortios_fortimanager_devicemanager_install_device"
sidebar_current: "docs-fortios-fortimanager-resource-devicemanager-install-device"
description: |-
  Provides a resource to install devicemanager script from FortiManager to the related device
---

# fortios_fortimanager_devicemanager_install_device
This resource supports installing devicemanager script from FortiManager to the related device

## Example Usage
```hcl
provider "fortios" {
	hostname = "192.168.88.100"
	username = "APIUser"
	passwd = "admin"
	product = "fortimanager"
	insecure = false
	cabundlefile = "/path/yourCA.crt"
}

resource "fortios_fortimanager_devicemanager_install_device" "test1" {
	target_devname = "FGVM64-test"
	timeout = 5
}
```

## Argument Reference
The following arguments are supported:

* `target_devname` - (Required) Target device name.
* `timeout` - Timeout for installing the script to the target, default: 3 minutes.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `target_devname` - Target device name.
* `timeout` - Timeout for installing the script to the target, default: 3 minutes.
