package services

import (
	"github.com/rackspace/gophercloud"
)


// List makes a request against the API to list hypervisors accessible to you
// in pagination way (Not supported now by OpenStack for hypervisors).
func List(client *gophercloud.ServiceClient) GetResult {
	var result GetResult
	_, result.Err = client.Get(getListURL(client),
		&result.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200, 203},
		})
	return result
}

func Disable(client *gophercloud.ServiceClient, binary string, host string) DisableResult {
	type request struct {
		Binary string
		Host string
	}

	req := request{
		Binary: binary,
		Host: host,
	}

	var result DisableResult
	_, result.Err = client.Request("PUT", disableURL(client), gophercloud.RequestOpts{
		JSONBody:     &req,
		JSONResponse: &result.Body,
		OkCodes:      []int{200},
	})
	return result
}

func DisableWithLog(client *gophercloud.ServiceClient, binary string, host string, log string) DisableWithLogResult {
	type request struct {
		Binary 	string
		Host 	string
		Log 	string	`json:"disabled_reason"`
	}

	req := request{
		Binary: binary,
		Host: host,
		Log: log,
	}

	var result DisableWithLogResult
	_, result.Err = client.Request("PUT", disableWithLogURL(client), gophercloud.RequestOpts{
		JSONBody:     &req,
		JSONResponse: &result.Body,
		OkCodes:      []int{200},
	})
	return result
}


func Enable(client *gophercloud.ServiceClient, binary string, host string) EnableResult {
	type request struct {
		Binary string
		Host string
	}

	req := request{
		Binary: binary,
		Host: host,
	}

	var result EnableResult
	_, result.Err = client.Request("PUT", enableURL(client), gophercloud.RequestOpts{
		JSONBody:     &req,
		JSONResponse: &result.Body,
		OkCodes:      []int{200},
	})
	return result
}

func ForceDown(client *gophercloud.ServiceClient, binary string, host string) ForceDownResult {
	type request struct {
		Binary string
		Host string
	}

	req := request{
		Binary: binary,
		Host: host,
	}

	var result ForceDownResult
	_, result.Err = client.Request("PUT", forceDownURL(client), gophercloud.RequestOpts{
		JSONBody:     &req,
		JSONResponse: &result.Body,
		OkCodes:      []int{200},
	})
	return result
}

func Delete(client *gophercloud.ServiceClient, id string) error {
	var result DeleteResult
	_, result.Err = client.Request("DELETE", deleteURL(client, id), gophercloud.RequestOpts{
		OkCodes:      []int{200},
	})
	return result.Err
}
