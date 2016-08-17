package services

import (
	"github.com/rackspace/gophercloud"
	"fmt"
)


func getListURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("os-services")
}

func disableURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("os-services/disable")
}

func disableWithLogURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("os-services/disable-log-reason")
}

func enableURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("os-services/enable")
}

func forceDownURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("os-services/force-down")
}

func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL(fmt.Sprintf("os-services/%s", id))
}