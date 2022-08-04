package orders

import "github.com/innodreamer/ktvpc-sdk_poc"

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("orders")
}

func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("orders", id)
}

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("orders")
}

func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("orders", id)
}