package receivers

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

var apiVersion = "v1"
var apiName = "receivers"

func commonURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL(apiVersion, apiName)
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return commonURL(client)
}

func idURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, apiName, id)
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return idURL(client, id)
}

func updateURL(client *ktvpcsdk.ServiceClient, id string) string {
	return idURL(client, id)
}

func listURL(client *ktvpcsdk.ServiceClient) string {
	return commonURL(client)
}

func deleteURL(client *ktvpcsdk.ServiceClient, id string) string {
	return idURL(client, id)
}

func notifyURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL(apiVersion, apiName, id, "notify")
}
