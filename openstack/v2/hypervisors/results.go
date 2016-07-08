package hypervisors

import (
	"github.com/rackspace/gophercloud"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"github.com/rackspace/gophercloud/pagination"
	"strconv"
	"fmt"
)

type Hypervisor struct {
	HypervisorHostname string	`mapstructure:"hypervisor_hostname"`
	Id int
	State string
	Status string
}

type HypervisorDetail struct {
	HostIP string 				`mapstructure:"host_ip"`
	ServiceHost string			`mapstructure:"service_host"`
	Id int
	RunningVMs int				`mapstructure:"running_vms"`
	FreeDiskGB int16			`mapstructure:"free_disk_gb"`
	HypervisorVersion int32		`mapstructure:"hypervisor_version"`
	DistAvailableLeast int16	`mapstructure:"disk_available_least"`
	LocalGB int16				`mapstructure:"local_gb"`
	FreeRamMB int32				`mapstructure:"free_ram_mb"`
	Status string
	VcpuUsed int16				`mapstructure:"vcpus_used"`
	HypervisorType string		`mapstructure:"hypervisor_type"`
	LocalGBUsed	int16			`mapstructure:"local_gb_used"`
	Vcpus int
	HypervisorHostname string	`mapstructure:"hypervisor_hostname"`
	MemoryMBUsed int32			`mapstructure:"memory_mb_used"`
	MemoryMB int32				`mapstructure:"memory_mb"`
	CurrentWorkload int			`mapstructure:"current_workload"`
	State string
}


// HypervisorPage abstracts the raw results of making a List() request against the API.
// As OpenStack extensions may freely alter the response bodies of structures returned to the client, you may only safely access the
// data provided through the ExtractServers call.
type HypervisorPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a page contains no Server results.
func (page HypervisorPage) IsEmpty() (bool, error) {
	hypervisors, err := ExtractHypervisors(page)
	if err != nil {
		return true, err
	}
	return len(hypervisors) == 0, nil
}

// NextPageURL uses the response's embedded link reference to navigate to the next page of results.
func (page HypervisorPage) NextPageURL() (string, error) {
	type resp struct {
		Links []gophercloud.Link
	}

	var r resp
	err := mapstructure.Decode(page.Body, &r)
	if err != nil {
		return "", err
	}

	return gophercloud.ExtractNextURL(r.Links)
}


// ExtractHypervisors interprets the results of a single page from a List() call, producing a slice of Server entities.
func ExtractHypervisors(page pagination.Page) ([]Hypervisor, error) {
	casted := page.(HypervisorPage).Body

	var response struct {
		Hypervisors []Hypervisor `mapstructure:"hypervisors"`
	}

	config := &mapstructure.DecoderConfig{
		DecodeHook: toMapFromString,
		Result:     &response,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(casted)

	return response.Hypervisors, err
}

type hypervisorResult struct {
	gophercloud.Result
}


// GetResult temporarily contains the response from a Get call.
type GetResult struct {
	hypervisorResult
}

// Extract interprets any hypervisorResult as a Hypervisor, if possible.
func (r hypervisorResult) Extract() (*HypervisorDetail, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Hypervisor HypervisorDetail `mapstructure:"hypervisor"`
	}

	config := &mapstructure.DecoderConfig{
		DecodeHook: toMapFromString,
		Result:     &response,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(r.Body)
	if err != nil {
		return nil, err
	}

	return &response.Hypervisor, nil
}


func toMapFromString(from reflect.Kind, to reflect.Kind, data interface{}) (interface{}, error) {
	if (from == reflect.String) && (to == reflect.Map) {
		return map[string]interface{}{}, nil
	}
	return data, nil
}


func ListHypervisorsDetails(client *gophercloud.ServiceClient) (hypervisorsDetailList []*HypervisorDetail){
	// Retrieve a pager (i.e. a paginated collection)
	pager := List(client)

	// Define an anonymous function to be executed on each page's iteration
	err := pager.EachPage(func(page pagination.Page) (bool, error) {
		hypervisorList, err := ExtractHypervisors(page)
		if err != nil {
			return false, err
		}
		for _, h := range hypervisorList {
			// "h" will be a hypervisors.Hypervisor
			details, err := GetDetail(client, strconv.Itoa(h.Id)).Extract()
			if err != nil {
				fmt.Println(err)
			}
			hypervisorsDetailList = append(hypervisorsDetailList, details)
		}
		return true, nil
	})
	if err != nil {
		fmt.Println(err)
	}
	return hypervisorsDetailList
}