// Package ping provides a quick solution for checking if a connection can be
// established to some address.
package ping

import (
	"net"
	"time"
)

const timeout = 500 * time.Millisecond

// TCP tries to connect to address over tcp. It returns true if a connection
// can be established within 500 ms.
func TCP(address string) bool {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
