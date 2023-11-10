package tcp

import (
	"fmt"
)

var (
	ContentSizeTooLargeErr = fmt.Errorf("content size too large")
	NilRequestErr          = fmt.Errorf("request is nil")
	InvalidProtocolErr     = fmt.Errorf("invalid protocol operation")
)
