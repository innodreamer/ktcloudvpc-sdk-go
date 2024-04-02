package policytypes

import "github.com/cloud-barista/ktcloudvpc-sdk-go"

const (
	apiVersion = "v1"
	apiName    = "policy-types"
)

func policyTypeListURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func policyTypeGetURL(client *gophercloud.ServiceClient, policyTypeName string) string {
	return client.ServiceURL(apiVersion, apiName, policyTypeName)
}
