package ping

import "testing"

func TestTCP(t *testing.T) {
	canConnect1 := TCP("www.google.com:1")
	canConnect2 := TCP("www.google.com:80")
	if canConnect1 {
		t.Error("Should not be able to connect!")
	}
	if !canConnect2 {
		t.Error("Should be able to connect!")
	}
}

func TestTLS(t *testing.T) {
	canConnect1 := TLS("www.google.com:1")
	canConnect2 := TLS("www.google.com:443")
	if canConnect1 {
		t.Error("Should not be able to connect!")
	}
	if !canConnect2 {
		t.Error("Should be able to connect!")
	}
}
