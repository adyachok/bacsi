package hypervisors

import (
	"github.com/rackspace/gophercloud"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"github.com/rackspace/gophercloud/pagination"
)


// Page is an abstract struct which beholds next page url logic.
// Pagination logic for every entity have to implement isEmpty function and
// include field of type Page.
// Openstack Nova Api allows to send next optional args in request:
//		limit - max quantity of entities to select
//		marker - last element viewed by a customer
//		page_size - quantity of entities on a page
// Assuming possibly big quantity of hypervisors to extract
// are implemented paging for hypervisors and hypervisors details
type Page struct {
	pagination.LinkedPageBase
}

// NextPageURL uses the response's embedded link reference to navigate to the
// next page of results.
func (page Page) NextPageURL() (string, error) {
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


type HypervisorPage struct {
	Page
}

// IsEmpty returns true if a page contains no Hypervisor results.
func (page HypervisorPage) IsEmpty() (bool, error) {
	hypervisors, err := ExtractHypervisors(page)
	if err != nil {
		return true, err
	}
	return len(hypervisors) == 0, nil
}


type HypervisorsDetailsPage struct {
	Page
}

// IsEmpty returns true if a page contains no Hypervisor results.
func (page HypervisorsDetailsPage) IsEmpty() (bool, error) {
	hypervisors, err := ExtractHypervisorsDetails(page)
	if err != nil {
		return true, err
	}
	return len(hypervisors) == 0, nil
}

// Decodes response body. Accepts empty entity pointer and initiates
// this entity with decoded values.
func processResponse(response interface{}, body interface{}) error{
	config := &mapstructure.DecoderConfig{
		DecodeHook: toMapFromString,
		Result:     response,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	err = decoder.Decode(body)
	if err != nil {
		return err
	}
	return nil
}

// ExtractHypervisors interprets the results of a single page from a List() call,
// producing a slice of Hypervisor entities.
func ExtractHypervisors(page pagination.Page) ([]Hypervisor, error) {
	casted := page.(HypervisorPage).Body

	var response struct {
		Hypervisors []Hypervisor `mapstructure:"hypervisors"`
	}
	err := processResponse(&response, casted)
	return response.Hypervisors, err
}

// ExtractHypervisorsDetails interprets the results of a single page from a List() call,
// producing a slice of HypervisorsDetails entities.
func ExtractHypervisorsDetails(page pagination.Page) ([]HypervisorDetail, error) {
	casted := page.(HypervisorsDetailsPage).Body

	var response struct {
		Hypervisors []HypervisorDetail `mapstructure:"hypervisors"`
	}
	err := processResponse(&response, casted)
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
func (r hypervisorResult) Extract() ([]Hypervisor, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Hypervisors []Hypervisor `mapstructure:"hypervisors"`
	}
	err := processResponse(&response, r.Body)
	return response.Hypervisors, err
}

// Extract interprets any hypervisorResult as a HypervisorDetails, if possible.
func (r hypervisorResult) ExtractDetails() ([]HypervisorDetail, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Hypervisor []HypervisorDetail `mapstructure:"hypervisors"`
	}
	err := processResponse(&response, r.Body)
	return response.Hypervisor, err
}

// Extract interprets any hypervisorResult as a HypervisorDetail, if possible.
func (r hypervisorResult) ExtractDetail() (*HypervisorDetail, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Hypervisor HypervisorDetail `mapstructure:"hypervisor"`
	}
	err := processResponse(&response, r.Body)
	return &response.Hypervisor, err
}

// Extract interprets any hypervisorResult as a HypervisorServersInfo, if possible.
func (r hypervisorResult) ExtractServersInfo() ([]HypervisorServersInfo, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Hypervisor []HypervisorServersInfo `mapstructure:"hypervisors"`
	}
	err := processResponse(&response, r.Body)
	return response.Hypervisor, err
}

// Extract interprets any hypervisorResult as a HypervisorUptime, if possible.
func (r hypervisorResult) ExtractUptime() (*HypervisorUptimeInfo, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Hypervisor HypervisorUptimeInfo `mapstructure:"hypervisor"`
	}
	err := processResponse(&response, r.Body)
	return &response.Hypervisor, err
}

func toMapFromString(from reflect.Kind, to reflect.Kind, data interface{}) (interface{}, error) {
	if (from == reflect.String) && (to == reflect.Map) {
		return map[string]interface{}{}, nil
	}
	return data, nil
}
