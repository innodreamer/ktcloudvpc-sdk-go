package testing

import (
	"net/http"
	"testing"

	th "github.com/cloud-barista/ktcloudvpc-sdk-go/testhelper"
	"github.com/cloud-barista/ktcloudvpc-sdk-go/testhelper/client"
)

func mockMigrateResponse(t *testing.T, id string) {
	th.Mux.HandleFunc("/servers/"+id+"/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{"migrate": null}`)
		w.WriteHeader(http.StatusAccepted)
	})
}

func mockLiveMigrateResponse(t *testing.T, id string) {
	th.Mux.HandleFunc("/servers/"+id+"/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
			"os-migrateLive": {
				"host": "01c0cadef72d47e28a672a76060d492c",
				"block_migration": false,
				"disk_over_commit": true
			}
		}`)
		w.WriteHeader(http.StatusAccepted)
	})
}
