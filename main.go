// main.go - Main file for the go-mac-converter program which is a golang version of the python mac converter program
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	//Variables
	var newMac string

	// Check if the mac address(s) are inputed on the commandline. If not check if a file name was given instead
	// Get the converted mac address(s) and print the mac address(s) out to the CLI
	macAddr, file, inputType, outputType := cli()

	if macAddr != "" {
		newMac = getMac(macAddr, inputType, outputType)
		printMac(newMac)
	} else if file != "" {
		fileName, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer fileName.Close()

		scanner := bufio.NewScanner(fileName)
		for scanner.Scan() {
			newMac = getMac(scanner.Text(), inputType, outputType)
			printMac(newMac)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
