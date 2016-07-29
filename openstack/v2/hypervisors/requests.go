package hypervisors

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
)

// TODO: introduce list options for the pagination - marker, limit, page_size
// as this is defined in gophercloud/openstack/v2/servers/requests.go:{row 25}

// List makes a request against the API to list hypervisors accessible to you
// in pagination way.
func ListPaginated(client *gophercloud.ServiceClient) pagination.Pager {
	url := getListURL(client)

	createPageFn := func(r pagination.PageResult) pagination.Page {
		return HypervisorPage{Page{pagination.LinkedPageBase{PageResult: r}}}
	}

	return pagination.NewPager(client, url, createPageFn)
}

// List makes a request against the API to list hypervisors accessible to you.
func List(client *gophercloud.ServiceClient) GetResult {
	var result GetResult
	_, result.Err = client.Get(getListURL(client),
		&result.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200, 203},
		})
	return result
}


// Detailed list makes a request against the API to list details of all
// hypervisors accessible to you.
func GetDetailesListPaginated(client *gophercloud.ServiceClient) pagination.Pager {
	url := getDetailedListURL(client)

	createPageFn := func(r pagination.PageResult) pagination.Page {
		return HypervisorsDetailsPage{Page{pagination.LinkedPageBase{PageResult: r}}}
	}

	return pagination.NewPager(client, url, createPageFn)
}

// Get details about specified by id hypervisor
func GetDetailsList(client *gophercloud.ServiceClient) GetResult {
	var result GetResult
	_, result.Err = client.Get(getDetailedListURL(client), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 203},
	})
	return result
}

// Get details about specified by id hypervisor
func GetDetail(client *gophercloud.ServiceClient, id string) GetResult {
	var result GetResult
	_, result.Err = client.Get(getDetailURL(client, id), &result.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200, 203},
	})
	return result
}

// Get VMs booted on specified by hostname hypervisor
func GetHypervisorServers(client *gophercloud.ServiceClient, hypervisorHostname string) GetResult {
	var result GetResult
	_, result.Err = client.Get(getHypervisorServersURL(client, hypervisorHostname),
		&result.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200, 203},
		})
	return result
}

func GetHypervisorUptime(client *gophercloud.ServiceClient, id string) GetResult {
	var result GetResult
	_, result.Err = client.Get(getHypervisorUptimeURL(client, id),
		&result.Body, &gophercloud.RequestOpts{
			OkCodes: []int{200, 203},
		})
	return result
}
