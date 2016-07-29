package main

import (
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
	"fmt"
	"os"
	"./openstack/v2/hypervisors"
	g_servers "./openstack/v2/servers"


)


func main() {
	serversMap := make(map[string][]*servers.Server)
	client := getOpenStackClient()
	details, _ := hypervisors.List(client).Extract()
	for _, h_details := range details {
		hostName := h_details.HypervisorHostname
		serversList, _ := g_servers.ListServersByHost(client, hostName)
		serversMap[hostName] = serversList
	}
	fmt.Println(serversMap)
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
