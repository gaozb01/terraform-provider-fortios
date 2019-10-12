package fmgclient

import (
	"fmt"
	"strconv"

	"github.com/fgtdev/fortimanager-sdk-go/util"
)

// JSONSystemAdom contains the params for creating system adom
type JSONSystemAdom struct {
	Name           string   `json:"name"`
	RestrictedPrds string   `json:"restricted_prds"`
	Status         string   `json:"state"`
	Flags          []string `json:"flags"`
}

// CreateUpdateSystemAdom is for creating or updating the system adom
func (c *FmgSDKClient) CreateUpdateSystemAdom(params *JSONSystemAdom, method string) (err error) {
	defer c.Trace("CreateUpdateSystemAdom")()

	p := map[string]interface{}{
		"data": params,
		"url":  "/dvmdb/adom",
	}

	_, err = c.Do(method, p)

	if err != nil {
		err = fmt.Errorf("CreateUpdateSystemAdom failed: %s", err)
		return
	}

	return
}

// ReadSystemAdom is for reading the specific system adom
func (c *FmgSDKClient) ReadSystemAdom(name string) (out *JSONSystemAdom, err error) {
	defer c.Trace("ReadSystemAdom")()

	p := map[string]interface{}{
		"url": "/dvmdb/adom/" + name,
	}

	result, err := c.Do("get", p)
	if err != nil {
		err = fmt.Errorf("ReadSystemAdom failed :%s", err)
		return
	}

	data := (result["result"].([]interface{}))[0].(map[string]interface{})["data"].(map[string]interface{})
	if data == nil {
		err = fmt.Errorf("cannot get the results from the response")
		return
	}

	out = &JSONSystemAdom{}
	if data["name"] != nil {
		out.Name = data["name"].(string)
	}
	if data["state"] != nil {
		out.Status = strconv.Itoa(int(data["state"].(float64)))
	}
	if data["restricted_prds"] != nil {
		out.RestrictedPrds = util.RestrictedPrds2Str(int(data["restricted_prds"].(float64)))
	}

	return

}

// DeleteSystemAdom is for deleting the specific system adom
func (c *FmgSDKClient) DeleteSystemAdom(name string) (err error) {
	defer c.Trace("DeleteSystemAdom")()

	p := map[string]interface{}{
		"url": "/dvmdb/adom/" + name,
	}

	_, err = c.Do("delete", p)
	if err != nil {
		err = fmt.Errorf("DeleteSystemAdom failed :%s", err)
		return
	}

	return
}
