package orders

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

func listURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("orders")
}

func getURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("orders", id)
}

func createURL(client *ktvpcsdk.ServiceClient) string {
	return client.ServiceURL("orders")
}

func deleteURL(client *ktvpcsdk.ServiceClient, id string) string {
	return client.ServiceURL("orders", id)
}
