package services

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud"
)

type commonResult struct {
	gophercloud.Result
}

// Extract interprets a GetResult, CreateResult or UpdateResult as a []Service.
// An error is returned if the original call or the extraction failed.
func (r commonResult) ExtractServices() ([]*Service, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var res struct {
		Services []*Service `json:"services"`
	}

	err := mapstructure.Decode(r.Body, &res)

	return res.Services, err
}

// Extract interprets a GetResult, CreateResult or UpdateResult as a concrete Service.
// An error is returned if the original call or the extraction failed.
func (r commonResult) Extract() (*Service, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var res struct {
		Service `json:"service"`
	}

	err := mapstructure.Decode(r.Body, &res)

	return &res.Service, err
}

// GetResult is the deferred result of a Get call.
type GetResult struct {
	commonResult
}

// DisableResult is the deferred result of an Update call.
type DisableResult struct {
	commonResult
}

// DisableWithLogResult is the deferred result of an Update call.
type DisableWithLogResult struct {
	commonResult
}

// EnableResult is the deferred result of an Update call.
type EnableResult struct {
	commonResult
}

// ForceDownResult is the deferred result of an Update call.
type ForceDownResult struct {
	commonResult
}

// DeleteResult is the deferred result of an Delete call.
type DeleteResult struct {
	gophercloud.ErrResult
}


type Service struct {
	Id            int
	Binary        string
	DisableReason string 	`mapstructure:"disabled_reason"`
	Host          string
	State         string
	Status        string
	UpdatedAt     string 	`mapstructure:"updated_at"`
	ForcedDown    bool	 	`mapstructure:"forced_down"`
	Zone          string
}
