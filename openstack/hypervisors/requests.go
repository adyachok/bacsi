package hypervisors

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)


func GetDetail(client *gophercloud.ServiceClient, id string) GetResult {
	var result GetResult
	_, result.Err = client.Get(getDetailURL(client, id), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 203},
	})
	return result
}

// List makes a request against the API to list hypervisors accessible to you.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	url := getListURL(client)

	createPageFn := func(r pagination.PageResult) pagination.Page {
		return HypervisorPage{pagination.LinkedPageBase{PageResult: r}}
	}

	return pagination.NewPager(client, url, createPageFn)
}