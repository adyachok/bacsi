package hypervisors

import "github.com/rackspace/gophercloud"


func getListURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("os-hypervisors")
}


func getDetailURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("os-hypervisors", id)
}