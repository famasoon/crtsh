package crtlog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// CRTSHURL is URL of crt.sh endpoint
const CRTSHURL string = "https://crt.sh/"

// CTLogs is CTLog slice
type CTLogs []*CTLog

// CTLog is result to query crt.sh
// CTLog include some information related CA
type CTLog struct {
	IssuerCaID        int    `json:"issuer_ca_id"`
	IssuerName        string `json:"issuer_name"`
	NameValue         string `json:"name_value"`
	MinCertID         int    `json:"min_cert_id"`
	MinEntryTimestamp string `json:"min_entry_timestamp"`
	NotBefore         string `json:"not_before"`
	NotAfter          string `json:"not_after"`
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
		err = fmt.Errorf("can not Access crt.sh")
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (ctlog CTLog) showFullCTlog() {
	fmt.Println("{")
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

// SearchComon query in crt.sh by common name and print result that
func SearchCommon(query string, onlyDomainFlag bool) error {
	var ctlogs CTLogs

	res, err := queryCrtsh(CRTSHURL + "?output=json&CN=" + query)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(res, &ctlogs); err != nil {
		return err
	}

	if onlyDomainFlag {
		for _, ctlog := range ctlogs {
			fmt.Printf("%s\n", ctlog.NameValue)
		}
	} else {
		for _, ctlog := range ctlogs {
			ctlog.showFullCTlog()
		}
	}

	return nil
}

// QueryCrt query in crt.sh and print result that
func QueryCrt(query string, onlyDomainFlag bool) error {
	var ctlogs CTLogs

	res, err := queryCrtsh(CRTSHURL + "?output=json&q=" + query)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(res, &ctlogs); err != nil {
		return err
	}

	if onlyDomainFlag {
		for _, ctlog := range ctlogs {
			fmt.Printf("%s\n", ctlog.NameValue)
		}
	} else {
		for _, ctlog := range ctlogs {
			ctlog.showFullCTlog()
		}
	}

	return nil
}

// GetPemFile download pemfile from crt.sh and dump to []byte
func GetPemFile(certID int) ([]byte, error) {
	body, err := queryCrtsh(CRTSHURL + "?d=" + strconv.Itoa(certID))
	if err != nil {
		return nil, err
	}

	return body, nil
}
