//
// goamz - Go packages to interact with the Amazon Web Services.
//
//   https://wiki.ubuntu.com/goamz
//
package sdb

import (
	"github.com/TelAPI/amz"
)

func Sign(auth aws.Auth, method, path string, params map[string][]string, headers map[string][]string) {
	sign(auth, method, path, params, headers)
}
