---
layout: "fortios"
page_title: "Provider: FortiOS"
sidebar_current: "docs-fortios-index"
description: |-
  The FortiOS provider interacts with FortiGate or FortiManager.
---

# FortiOS Provider

The FortiOS provider consists of two products: FortiGate and FortiManager, it can configure and manager them by the related resources.

For the usage and details about the FortiGate and FortiManager resources, please refer to the related documents.

## Argument Reference for FortiGate
The following arguments are supported:

* `hostname` - (Required) Hostname or IP address of FortiGate.
* `token` - (Required) Token of FortiGate.
* `insecure` - (Optional) Control if the provider performs insecure SSL request, default is `false`.
* `cabundlefile` - (Optional) The path of a custom CA bundle file.
* `vdom` - (Optional) Specify the vdom name if the FortiGate is running in VDOM mode.

As the alternative choice, the following environment variables are also supported:

* FORTIOS_ACCESS_HOSTNAME
* FORTIOS_ACCESS_TOKEN
* FORTIOS_INSECURE
* FORTIOS_CA_CABUNDLE

For `token` argument, please refer to the `system->system api-user` and `execute->api-user` chapters of the `FortiOS Handbook - CLI Reference` for more details.

### Example Usage

```hcl
provider "fortios" {
	hostname = "192.168.52.177"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
	insecure = "false"
	cabundlefile = "/path/yourCA.crt"
}

# Create a Static Route Item
resource "fortios_networking_route_static" "test1" {
	dst = "110.2.2.122/32"
	gateway = "2.2.2.2"
	# ...
}
```

Note that `insecure` could be set to `true` for a quick test, but is not recommended for a real usage.

```hcl
provider "fortios" {
	hostname = "192.168.52.177"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
	insecure = "true"
}
```

## Argument Reference for FortiManager
The following arguments are supported:

* `hostname` - (Required) Hostname or IP address of FortiManager.
* `username` - (Required) Admin username for remote API communication with FortiManager.
* `passwd` - (Required) Password.
* `product` - (Required) Which product to use, should be set to "fortimanager" here.
* `insecure` - (Optional) Control if the provider performs insecure SSL request, default is `false`.
* `cabundlefile` - (Optional) The path of a custom CA bundle file.

As the alternative choice, the following environment variables are also supported:

* FORTIOS_FMG_HOSTNAME
* FORTIOS_FMG_USERNAME
* FORTIOS_FMG_PASSWORD
* FORTIOS_PRODUCT
* FORTIOS_FMG_INSECURE
* FORTIOS_FMG_CABUNDLE

### Example Usage

```hcl
provider "fortios" {
        hostname = "192.168.88.100"
        username = "APIUser"
        passwd = "admin"
        product = "fortimanager"
        insecure = false
        cabundlefile = "/path/yourCA.crt"
}

resource "fortios_fortimanager_system_dns" "test1" {
        primary = "208.91.112.52"
        secondary = "208.91.112.54"
}
```

Note that `insecure` could be set to `true` for a quick test, but is not recommended for a real usage.

## Versioning

The provider can cover both v6.0 and v6.2 versions.
