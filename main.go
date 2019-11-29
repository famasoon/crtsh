package main

import (
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/famasoon/crtsh/crtlog"
	"github.com/famasoon/crtsh/parser"
)

func showUsage() {
	fmt.Println("This tool shows the result of crt.sh")
	fmt.Println("Option:")
	fmt.Println("  -q Query")
	fmt.Println("  -i Min Cert ID")
	fmt.Println("  -cn Common Name")
	fmt.Printf("Usage: %s -q example.com\n", os.Args[0])
	os.Exit(0)
}

func parseCTLog(certID int) error {
	body, err := crtlog.GetPemFile(certID)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(body)

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}

	fmt.Println("Enumrate DNS Names:")
	for _, dnsName := range parser.EnumDNS(cert) {
		fmt.Println(dnsName)
	}
	return nil
}

// TODO: Create run function () (err)
func main() {
	var (
		query          string
		certID         int
		onlyDomainFlag bool
		commonName     string
	)
	flag.StringVar(&query, "q", "", "Query String")
	flag.BoolVar(&onlyDomainFlag, "o", false, "Print only domains")
	flag.IntVar(&certID, "i", 0, "Min Cert ID")
	flag.StringVar(&commonName, "cn", "", "Query string for common name")
	flag.Parse()
	if query == "" && certID == 0 && commonName == "" {
		showUsage()
	}

	if query != "" {
		err := crtlog.QueryCrt(query, onlyDomainFlag)
		if err != nil {
			log.Fatal(err)
		}
	} else if certID != 0 {
		fmt.Printf("CertID: %d\n", certID)

		err := parseCTLog(certID)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := crtlog.SearchComon(commonName, onlyDomainFlag)
		if err != nil {
			log.Fatal(err)
		}
	}
}
