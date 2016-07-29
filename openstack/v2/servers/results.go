package servers

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
)


func ListServers(client *gophercloud.ServiceClient) (serversList []*servers.Server, error){
	// We have the option of filtering the server list. If we want the full
	// collection, leave it as an empty struct
	opts := servers.ListOpts{AllTenants: true}

	// Retrieve a pager (i.e. a paginated collection)
	pager := servers.List(client, opts)

	// Define an anonymous function to be executed on each page's iteration
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		serverList, err := servers.ExtractServers(page)
		if err != nil {
			return false, err
		}
		for _, s := range serverList {
			// "s" will be a servers.Server
			serversList = append(serversList, s)

		}
		return true, nil
	})
	return &serversList, err
}

// ListServersByHost create a list of instances booted on a host
func ListServersByHost(client *gophercloud.ServiceClient, hostName string) (serversList []*servers.Server, error) {
	// We have the option of filtering the server list. If we want the full
	// collection, leave it as an empty struct
	opts := servers.ListOpts{Host: hostName, AllTenants: true}

	// Retrieve a pager (i.e. a paginated collection)
	pager := servers.List(client, opts)

	// Define an anonymous function to be executed on each page's iteration
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		serverList, err := servers.ExtractServers(page)
		if err != nil {
			return false, err
		}
		for _, s := range serverList {
			// "s" will be a servers.Server
			serversList = append(serversList, &s)

		}
		return true, nil
	})
	return &serversList, err
}
