package services

import (
	"github.com/rackspace/gophercloud"
	"fmt"
)

/*
 GET	/os-services							List Compute Services
 PUT	/os-services/disable					Disable Scheduling For A Compute Service
 PUT	/os-services/disable-log-reason			Log	Disabled Compute Service Information
 PUT	/os-services/enable						Enable Scheduling For A Compute Service
 PUT	/os-services/force-down					Update Forced Down
 DELETE	/os-services/{service_id}				Delete Compute Service
*/

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