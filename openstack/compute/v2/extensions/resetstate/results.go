package resetstate

import (
	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

// ResetResult is the response of a ResetState operation. Call its ExtractErr
// method to determine if the request suceeded or failed.
type ResetResult struct {
	ktvpcsdk.ErrResult
}
