package services

import (
	"testing"

	th "github.com/rackspace/gophercloud/testhelper"
	"github.com/rackspace/gophercloud/testhelper/client"
)


func TestListServices(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleServicesGetSuccessfully(t)

	actual, err := List(client.ServiceClient()).ExtractServices()

	th.AssertNoErr(t, err)

	if len(actual) != 4 {
		t.Errorf("Expected 4 services, saw %d", len(actual))
	}

	th.CheckDeepEquals(t, &ServicesExpected[0], actual[0])
	th.CheckDeepEquals(t, &ServicesExpected[1], actual[1])
	th.CheckDeepEquals(t, &ServicesExpected[2], actual[2])
	th.CheckDeepEquals(t, &ServicesExpected[3], actual[3])
}

func TestDisableService(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleServiceDisableSuccessfully(t)

	actual, err := Disable(client.ServiceClient(), "nova-compute", "host1").Extract()

	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, &NovaDisableExpected, actual)
}

func TestDisableWithLogService(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleServiceDisableWithLogSuccessfully(t)

	actual, err := DisableWithLog(client.ServiceClient(), "nova-compute",
		"host1", "test reason").Extract()

	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, &NovaDisableWithLogExpected, actual)
}

func TestEnableService(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleServiceEnableSuccessfully(t)

	actual, err := Enable(client.ServiceClient(), "nova-compute", "host1").Extract()

	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, &NovaEnableExpected, actual)
}

func TestForceDownService(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleServiceForceDownSuccessfully(t)

	actual, err := ForceDown(client.ServiceClient(), "nova-compute", "host1").Extract()

	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, &NovaForceDownExpected, actual)
}

func TestDeleteService(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleServiceDeleteSuccessfully(t)

	err := Delete(client.ServiceClient(), "1")

	th.AssertNoErr(t, err)
}
