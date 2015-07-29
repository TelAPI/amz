//
// goamz - Go packages to interact with the Amazon Web Services.
//
//   https://wiki.ubuntu.com/goamz
//
package mturk

import (
	"github.com/TelAPI/amz"
)

func Sign(auth aws.Auth, service, method, timestamp string, params map[string]string) {
	sign(auth, service, method, timestamp, params)
}
