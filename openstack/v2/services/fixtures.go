package services

import (
	"net/http"
	"fmt"
	"testing"

	th "github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
)

const getListBody = `
{
    "services": [
        {
            "id": 1,
            "binary": "nova-scheduler",
            "disabled_reason": "test1",
            "host": "host1",
            "state": "up",
            "status": "disabled",
            "updated_at": "2012-10-29T13:42:02.000000",
            "forced_down": false,
            "zone": "internal"
        },
        {
            "id": 2,
            "binary": "nova-compute",
            "disabled_reason": "test2",
            "host": "host1",
            "state": "up",
            "status": "disabled",
            "updated_at": "2012-10-29T13:42:05.000000",
            "forced_down": false,
            "zone": "nova"
        },
        {
            "id": 3,
            "binary": "nova-scheduler",
            "disabled_reason": null,
            "host": "host2",
            "state": "down",
            "status": "enabled",
            "updated_at": "2012-09-19T06:55:34.000000",
            "forced_down": false,
            "zone": "internal"
        },
        {
            "id": 4,
            "binary": "nova-compute",
            "disabled_reason": "test4",
            "host": "host2",
            "state": "down",
            "status": "disabled",
            "updated_at": "2012-09-18T08:03:38.000000",
            "forced_down": false,
            "zone": "nova"
        }
    ]
}
`
var (
	NovaSchedulerHost1Expected = Service{
		Id: 1,
		Binary: "nova-scheduler",
		DisableReason: "test1",
		Host: "host1",
		State: "up",
		Status: "disabled",
		UpdatedAt: "2012-10-29T13:42:02.000000",
		ForcedDown: false,
		Zone: "internal",

	}
	NovaComputeHost1Expected = Service{
		Id: 2,
		Binary: "nova-compute",
		DisableReason: "test2",
		Host: "host1",
		State: "up",
		Status: "disabled",
		UpdatedAt: "2012-10-29T13:42:05.000000",
		ForcedDown: false,
		Zone: "nova",
	}
	NovaSchedulerHost2Expected = Service{
		Id: 3,
		Binary: "nova-scheduler",
		DisableReason: "",
		Host: "host2",
		State: "down",
		Status: "enabled",
		UpdatedAt: "2012-09-19T06:55:34.000000",
		ForcedDown: false,
		Zone: "internal",
	}
	NovaComputeHost2Expected = Service{
		Id: 4,
		Binary: "nova-compute",
		DisableReason: "test4",
		Host: "host2",
		State: "down",
		Status: "disabled",
		UpdatedAt: "2012-09-18T08:03:38.000000",
		ForcedDown: false,
		Zone: "nova",
	}
	ServicesExpected = []Service{NovaSchedulerHost1Expected, NovaComputeHost1Expected,
		NovaSchedulerHost2Expected, NovaComputeHost2Expected}
)

func HandleServicesGetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-services", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, getListBody)
	})
}

const DisableBody = `
{
    "service": {
        "binary": "nova-compute",
        "host": "host1",
        "status": "disabled"
    }
}
`

var NovaDisableExpected = Service{
	Binary: "nova-compute",
	Host: "host1",
	Status: "disabled",
}

func HandleServiceDisableSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-services/disable", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, DisableBody)
	})
}

const DisableWithLogBody = `
{
    "service": {
        "binary": "nova-compute",
        "disabled_reason": "test reason",
        "host": "host1",
        "status": "disabled"
    }
}
`

var NovaDisableWithLogExpected = Service{
	Binary: "nova-compute",
	Host: "host1",
	Status: "disabled",
	DisableReason: "test reason",
}

func HandleServiceDisableWithLogSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-services/disable-log-reason", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, DisableWithLogBody)
	})
}

const EnableBody = `
{
    "service": {
        "binary": "nova-compute",
        "host": "host1",
        "status": "enabled"
    }
}
`

var NovaEnableExpected = Service{
	Binary: "nova-compute",
	Host: "host1",
	Status: "enabled",
}

func HandleServiceEnableSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-services/enable", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, EnableBody)
	})
}

const ForceDownBody = `
{
    "service": {
        "binary": "nova-compute",
        "host": "host1",
        "forced_down": true
    }
}
`

var NovaForceDownExpected = Service{
	Binary: "nova-compute",
	Host: "host1",
	ForcedDown: true,
}

func HandleServiceForceDownSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-services/force-down", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ForceDownBody)
	})
}

func HandleServiceDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/os-services/1", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
	})
}