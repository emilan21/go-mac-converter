package main

import (
	"fmt"
	"strings"
	"flag"
	"os"
	"log"
)

// Normalize the mac address and then convert the make address.
func getMac(mac, inputType, outputType string) string {
	normalizeMac := normalize(mac, inputType)
	convertedMac := convertMac(normalizeMac, inputType, outputType)
	return convertedMac
}

// Normalize the mac address to colon between every two digits based on the input type
func normalize(mac, inputType string) string {
	var normalizeMac string
	if inputType == "colon" {
		normalizeMac = mac
	}

	if inputType == "hp" {
		s := strings.Split(mac, "-")
		subString1, subString2, subString3 := s[0], s[1], s[2]
		normalizeMac = subString1[0:2] + ":" + subString1[2:4] + ":" + subString2[0:2] + ":" + subString2[2:4] + ":" + subString3[0:2] + ":" + subString3[2:4]

	}

	if inputType == "no-delimiter" {
		normalizeMac = mac[0:2] + ":" + mac[2:4] + ":" + mac[4:6] + ":" + mac[6:8] + ":" + mac[8:10] + ":" + mac[10:12]
	}

	if inputType == "dash" {
		normalizeMac = strings.Replace(mac, "-", ":", -1)
	}

	return normalizeMac
}

// Convert the mac address from the normalized mac address to the correct output type
func convertMac(normalizeMac, inputType, outputType string) string {
	var convertedMac string
	if outputType == "colon" {
		convertedMac = normalizeMac
	}

	if outputType == "hp" {
		s := strings.Split(normalizeMac, ":")
		subString1, subString2, subString3, subString4, subString5, subString6 := s[0], s[1], s[2], s[3], s[4], s[5]
		convertedMac = subString1 + subString2 + "-" + subString3 + subString4 + "-" + subString5 + subString6
	}

	if outputType == "no-delimiter" {
		s := strings.Split(normalizeMac, ":")
		subString1, subString2, subString3, subString4, subString5, subString6 := s[0], s[1], s[2], s[3], s[4], s[5]
		convertedMac = subString1 + subString2 + subString3 + subString4 + subString5 + subString6
	}

	if outputType == "dash" {
		convertedMac = strings.Replace(normalizeMac, ":", "-", -1)
	}

	return convertedMac
}

// Print out mac address in it's correct output type
func printMac(newMac string) {
	fmt.Printf("%s\n", newMac)
}

func cli() (string, string, string, string) {
	// Command line flags handled
	macAddr := flag.String("macs", "", "Enter one or more mac addresses")
	file := flag.String("file", "", "File containing mac addresses")
	inputType := flag.String("input-type", "",
		`Type of mac address notation the input is in.
	
	 The mac address types are:
			 - colon
			 - hp
			 - no_delimiter
			 - dash`)
	outputType := flag.String("output-type", "",
		`Type of mac address notation you want the mac address in.
	
	 The mac address types are:
			 - colon
			 - hp
			 - no_delimiter
			 - dash`)

	// Parse command line flags
	flag.Parse()

	// If no command line arguments print helpful message then return. If not input type and/or output type entered on the command line.
	// Also print a helpful message then return.
	if len(os.Args) < 1 {
		log.Fatal("No command line arguments. Type -h for help")
	}
	if *inputType == "" {
		log.Fatal("No input type entered. Input type needed. Type -h for help")
	}
	if *outputType == "" {
		log.Fatal("No output type entered. Output type needed. Type -h for help")
	}

	return *macAddr, *file, *inputType, *outputType
}
