package main

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/famasoon/crtsh/crtlog"
	"github.com/famasoon/crtsh/parser"
)

// CRTSHURL is URL of crt.sh endpoint
const CRTSHURL string = "https://crt.sh/"

func showUsage() {
	fmt.Println("This tool shows the result of crt.sh")
	fmt.Println("Option:")
	fmt.Println("  -q Query")
	fmt.Println("  -i Min Cert ID")
	fmt.Println("  -cn Common Name")
	fmt.Printf("Usage: %s -q example.com\n", os.Args[0])
	os.Exit(0)
}

func queryCrtsh(query string) ([]byte, error) {
	req := http.Client{
		Timeout: 60 * time.Second,
	}
	res, err := req.Get(query)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = fmt.Errorf("Can not Access crt.sh")
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func searchComon(query string, onlyDomainFlag bool) error {
	var ctlogs crtlog.CTLogs

	res, err := queryCrtsh(CRTSHURL + "?output=json&CN=" + query)
	if err = json.Unmarshal(res, &ctlogs); err != nil {
		return err
	}

	if onlyDomainFlag {
		for _, ctlog := range ctlogs {
			fmt.Printf("%s\n", ctlog.NameValue)
		}
	} else {
		for key, ctlog := range ctlogs {
			fmt.Println("{")
			fmt.Printf("  Index: %d\n", key+1)
			fmt.Printf("  Issuer CA ID: %d\n", ctlog.IssuerCaID)
			fmt.Printf("  Issuer Name: %s\n", ctlog.IssuerName)
			fmt.Printf("  Name: %s\n", ctlog.NameValue)
			fmt.Printf("  Min Cert ID: %d\n", ctlog.MinCertID)
			fmt.Printf("  Min Entry TimeStamp: %s\n", ctlog.MinEntryTimestamp)
			fmt.Printf("  Not Before: %s\n", ctlog.NotBefore)
			fmt.Printf("  Not After: %s\n", ctlog.NotAfter)
			fmt.Printf("  Donwload Pem file: %s?d=%d\n", CRTSHURL, ctlog.MinCertID)
			fmt.Println("}")
		}
	}

	return nil
}

func queryCrt(query string, onlyDomainFlag bool) error {
	var ctlogs crtlog.CTLogs

	res, err := queryCrtsh(CRTSHURL + "?output=json&q=" + query)
	if err = json.Unmarshal(res, &ctlogs); err != nil {
		return err
	}

	if onlyDomainFlag {
		for _, ctlog := range ctlogs {
			fmt.Printf("%s\n", ctlog.NameValue)
		}
	} else {
		for key, ctlog := range ctlogs {
			fmt.Println("{")
			fmt.Printf("  Index: %d\n", key+1)
			fmt.Printf("  Issuer CA ID: %d\n", ctlog.IssuerCaID)
			fmt.Printf("  Issuer Name: %s\n", ctlog.IssuerName)
			fmt.Printf("  Name: %s\n", ctlog.NameValue)
			fmt.Printf("  Min Cert ID: %d\n", ctlog.MinCertID)
			fmt.Printf("  Min Entry TimeStamp: %s\n", ctlog.MinEntryTimestamp)
			fmt.Printf("  Not Before: %s\n", ctlog.NotBefore)
			fmt.Printf("  Not After: %s\n", ctlog.NotAfter)
			fmt.Printf("  Donwload Pem file: %s?d=%d\n", CRTSHURL, ctlog.MinCertID)
			fmt.Println("}")
		}
	}

	return nil
}

func parseCTLog(certID int) error {
	res, err := http.Get(CRTSHURL + "?d=" + strconv.Itoa(certID))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = fmt.Errorf("Can not Access crt.sh")
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
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
		err := queryCrt(query, onlyDomainFlag)
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
		err := searchComon(commonName, onlyDomainFlag)
		if err != nil {
			log.Fatal(err)
		}
	}
}
