package listeners

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

const (
	rootPath       = "lbaas"
	resourcePath   = "listeners"
	statisticsPath = "stats"
)

func rootURL(c *ktvpcsdk.ServiceClient) string {
	return c.ServiceURL(rootPath, resourcePath)
}

func resourceURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id)
}

func statisticsRootURL(c *ktvpcsdk.ServiceClient, id string) string {
	return c.ServiceURL(rootPath, resourcePath, id, statisticsPath)
}
