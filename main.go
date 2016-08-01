package main

import (
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud"
	"fmt"
	"os"
	"./openstack/v2/hypervisors"
)

// Example how to create admin client and use it to query
// Nova about hypervisors info
func main() {
	client := getOpenStackClient()
	list, _ := hypervisors.List(client).Extract()
	for _, hyp := range list {
		fmt.Println(hyp.HypervisorHostname)
	}

}

func getOpenStackClient() *gophercloud.ServiceClient{
	opts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		fmt.Println("Could not extract OpenStack options from Env")
	}
	provider, err := openstack.AuthenticatedClient(opts)

	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
		Availability: gophercloud.AvailabilityAdmin,
	})
	return client
}
