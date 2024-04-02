package job

import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"

// Querying task statuses URL
func jobURL(sc *gophercloud.ServiceClient, jobId string) string {  			// Added
	return sc.ServiceURL("Etc?command=queryAsyncJob&jobid=", jobId)
}
