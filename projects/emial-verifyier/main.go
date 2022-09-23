package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Domain, hasMX, hasSPF, spfRecords, hasDMARC, dmarcRecord\n")

	for scanner.Scan() {
		checkMail(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Could not read from input: ", err)
	}
}

func checkMail(domain string) {
	var hasMX, hasSPF, hasDMARC bool

	var spfRecords, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Fatal(err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	textRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range textRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecords = record
		}
	}

	textDmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range textDmarcRecords {
		if strings.HasPrefix(record, "v=dmark1") {
			hasDMARC = true
			dmarcRecord = record
		}
	}

	log.Printf("%v %v %v %v %v %v", domain, hasMX, hasSPF, spfRecords, hasDMARC, dmarcRecord)
}
