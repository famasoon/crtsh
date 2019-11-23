package parser

import (
	"crypto/x509"
)

// EnumDNS return []string that contain Enumrated DNS names
func EnumDNS(cert *x509.Certificate) (DNSrecords []string) {
	return cert.DNSNames
}
