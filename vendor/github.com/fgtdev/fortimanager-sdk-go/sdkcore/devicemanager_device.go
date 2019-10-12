package fmgclient

import (
	"fmt"
)

// JSONDVMDeviceCreate is for device create
type JSONDVMDeviceCreate struct {
	UserId   string `json:"adm_usr"`
	Passwd   string `json:"adm_pass"`
	Ipaddr   string `json:"ip"`
	Name     string `json:"name"`
	MgmtMode string `json:"mgmt_mode"`
}

// JSONDVMDeviceDel is for device delete
type JSONDVMDeviceDel struct {
	Name string `json:"device"`
	Adom string `json:"adom"`
}

// CreateDVMDevice add device to devicemanager
func (c *FmgSDKClient) CreateDVMDevice(params *JSONDVMDeviceCreate) (err error) {
	defer c.Trace("CreateDVMDevice")()

	data := map[string]interface{}{
		"adom":   "root",
		"device": params,
	}

	p := map[string]interface{}{
		"data": data,
		"url":  "/dvm/cmd/add/device",
	}

	_, err = c.Do("exec", p)

	if err != nil {
		return fmt.Errorf("CreateDVMDevice failed: %s", err)
	}

	return
}

// DeleteDVMDevice del device from devicemanager
func (c *FmgSDKClient) DeleteDVMDevice(params *JSONDVMDeviceDel) (err error) {
	defer c.Trace("DeleteDVMDevice")()

	p := map[string]interface{}{
		"data": params,
		"url":  "/dvm/cmd/del/device",
	}

	_, err = c.Do("exec", p)

	if err != nil {
		return fmt.Errorf("DeleteDVMDevice failed: %s", err)
	}

	return
}
