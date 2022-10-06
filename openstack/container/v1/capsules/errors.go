package capsules

import (
	"fmt"

	"github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
)

type ErrInvalidDataFormat struct {
	ktvpcsdk.BaseError
}

func (e ErrInvalidDataFormat) Error() string {
	return fmt.Sprintf("Data in neither json nor yaml format.")
}
