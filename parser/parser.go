package parser

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/famasoon/crtsh/crtlog"
)

func ParseCTLog(certID int) error {
	body, err := crtlog.GetPemFile(certID)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(body)

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}

	fmt.Println("Enumerate DNS Names:")
	for _, dnsName := range enumDNS(cert) {
		fmt.Println(dnsName)
	}
	return nil
}

// EnumDNS return []string that contain Enumrated DNS names
func enumDNS(cert *x509.Certificate) (DNSrecords []string) {
	return cert.DNSNames
}
