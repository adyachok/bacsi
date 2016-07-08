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
	/*
			Using hrafd as a key-value store and serf as a gossip agent
		create distributed service which will track changes in the cluster
		(hosts failing and hosts restart) and will care about VMs plased on these
		hosts (evacuate and restart them).

			For this purpose data about hosts and VMs will be stored in key-value
		store and will be updated periodically.

		TODO: Host monitoring
			TODO: periodically get information about cluster
			TODO: in case of failed nodes evacuate VMs of these nodes
			TODO: What do in split cluster majority of CMHA nodes + minority of computes on one side?
		TODO: VM evacuation (depends on policy)
		TODO: Fencing and API for fencing

		TODO: почати з чогось простого
		 */

	serversMap := make(map[string][]*servers.Server)
	client := getOpenStackClient()
	details := hypervisors.ListHypervisorsDetails(client)
	for _, h_details := range details {
		hostName := h_details.HypervisorHostname
		serversList := g_servers.ListServersByHost(client, hostName)
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



