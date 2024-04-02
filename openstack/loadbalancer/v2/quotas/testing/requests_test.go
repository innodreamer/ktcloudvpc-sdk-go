package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/cloud-barista/ktcloudvpc-sdk-go"
	"github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/loadbalancer/v2/quotas"
	fake "github.com/cloud-barista/ktcloudvpc-sdk-go/openstack/networking/v2/common"
	th "github.com/cloud-barista/ktcloudvpc-sdk-go/testhelper"
)

func TestGet_1(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/quotas/0a73845280574ad389c292f6a74afa76", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponseRaw_1)
	})

	q, err := quotas.Get(fake.ServiceClient(), "0a73845280574ad389c292f6a74afa76").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, q, &GetResponse)
}

func TestGet_2(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/quotas/0a73845280574ad389c292f6a74afa76", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponseRaw_2)
	})

	q, err := quotas.Get(fake.ServiceClient(), "0a73845280574ad389c292f6a74afa76").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, q, &GetResponse)
}

func TestUpdate_1(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/quotas/0a73845280574ad389c292f6a74afa76", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, UpdateRequestResponseRaw_1)
	})

	q, err := quotas.Update(fake.ServiceClient(), "0a73845280574ad389c292f6a74afa76", quotas.UpdateOpts{
		Loadbalancer:  gophercloud.IntToPointer(20),
		Listener:      gophercloud.IntToPointer(40),
		Member:        gophercloud.IntToPointer(200),
		Pool:          gophercloud.IntToPointer(20),
		Healthmonitor: gophercloud.IntToPointer(-1),
		L7Policy:      gophercloud.IntToPointer(50),
		L7Rule:        gophercloud.IntToPointer(100),
	}).Extract()

	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, q, &UpdateResponse)
}

func TestUpdate_2(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/quotas/0a73845280574ad389c292f6a74afa76", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, UpdateRequestResponseRaw_2)
	})

	q, err := quotas.Update(fake.ServiceClient(), "0a73845280574ad389c292f6a74afa76", quotas.UpdateOpts{
		Loadbalancer:  gophercloud.IntToPointer(20),
		Listener:      gophercloud.IntToPointer(40),
		Member:        gophercloud.IntToPointer(200),
		Pool:          gophercloud.IntToPointer(20),
		Healthmonitor: gophercloud.IntToPointer(-1),
		L7Policy:      gophercloud.IntToPointer(50),
		L7Rule:        gophercloud.IntToPointer(100),
	}).Extract()

	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, q, &UpdateResponse)
}
