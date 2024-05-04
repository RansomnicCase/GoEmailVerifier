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
	fmt.Printf("domain, hasMX, hasSPF, spfRecords, hasDMARC, dmarcRecords\n")
	for scanner.Scan() {
		DomainChecker(scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		log.Printf("could not read the email:%v\n", err)
	}
}

func DomainChecker(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecords, dmarcRecord string

	MXrecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("an error occured: %v\n", err)
	}

	if len(MXrecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("error:%v\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecords = record
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("error:%v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecords, hasDMARC, dmarcRecord)
}
