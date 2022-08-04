//go:build acceptance || compute || rescueunrescue
// +build acceptance compute rescueunrescue

package v2

import (
	"testing"

	"github.com/innodreamer/ktvpc-sdk_poc/acceptance/clients"
	th "github.com/innodreamer/ktvpc-sdk_poc/testhelper"
)

func TestServerRescueUnrescue(t *testing.T) {
	client, err := clients.NewComputeV2Client()
	th.AssertNoErr(t, err)

	server, err := CreateServer(t, client)
	th.AssertNoErr(t, err)
	defer DeleteServer(t, client, server)

	err = RescueServer(t, client, server)
	th.AssertNoErr(t, err)

	err = UnrescueServer(t, client, server)
	th.AssertNoErr(t, err)
}
