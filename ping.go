// Package ping provides a quick solution for checking if a connection can be
// established to some address.
package ping

import (
	"crypto/tls"
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

// TLS tries to connect to address over tcp. It returns true if a connection
// can be established within 500 ms. Use this function instead of TCP() to
// avoid TLS handshake errors.
func TLS(address string) bool {
	d := &net.Dialer{Timeout: timeout}
	conn, err := tls.DialWithDialer(d, "tcp", address, nil)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
