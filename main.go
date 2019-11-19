package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/famasoon/crtsh/crtlog"
)

const CRTSH_URL = "https://crt.sh/"

// TODO: Create run function () (err)
func main() {
	var query string
	flag.StringVar(&query, "q", "", "Query String")
	flag.Parse()
	// TODO: Make Show Usage function
	if query == "" {
		fmt.Println("This tool shows the result of crt.sh")
		fmt.Println("Option:")
		fmt.Println("  -q Query")
		fmt.Printf("Usage: %s -q example.com\n", os.Args[0])
		os.Exit(0)
	}

	fmt.Println(query)

	// TODO: Create HTTP Client
	res, err := http.Get(CRTSH_URL + "?output=json&q=" + query)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = fmt.Errorf("Can not Access crt.sh")
		log.Fatal(err)
	}

	var ctlogs crtlog.CTLogs

	if err = json.NewDecoder(res.Body).Decode(&ctlogs); err != nil {
		log.Fatal(err)
	}

	for key, ctlog := range ctlogs {
		fmt.Printf("Index: %d\n", key+1)
		fmt.Printf("Issuer CA ID: %d\n", ctlog.IssuerCaID)
		fmt.Printf("Issuer Name: %s\n", ctlog.IssuerName)
		fmt.Printf("Name: %s\n", ctlog.NameValue)
		fmt.Printf("Min Cert ID: %d\n", ctlog.MinCertID)
		fmt.Printf("Min Entry TimeStamp: %s\n", ctlog.MinEntryTimestamp)
		fmt.Printf("Not Before: %s\n", ctlog.NotBefore)
		fmt.Printf("Not After: %s\n", ctlog.NotAfter)
		fmt.Println()
	}
}
