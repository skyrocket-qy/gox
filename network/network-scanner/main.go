package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/skyrocket-qy/gox/network/network-scanner/utils"
)

func main() {
	var (
		hostname, protocol string
		startPort, endPort int
	)

	log.Print("Enter hostname or IP: ")

	if _, err := fmt.Scanln(&hostname); err != nil {
		log.Fatalf("Error reading hostname: %v", err)
	}

	log.Print("Enter protocol (tcp or tcp6): ")

	if _, err := fmt.Scanln(&protocol); err != nil {
		log.Fatalf("Error reading protocol: %v", err)
	}

	log.Print("Enter start port: ")

	if _, err := fmt.Scanln(&startPort); err != nil {
		log.Fatalf("Error reading start port: %v", err)
	}

	log.Print("Enter end port: ")

	if _, err := fmt.Scanln(&endPort); err != nil {
		log.Fatalf("Error reading end port: %v", err)
	}

	log.Printf(
		"Scanning %s (%s) for open ports (%d-%d)...\n",
		hostname,
		protocol,
		startPort,
		endPort,
	)

	var wg sync.WaitGroup

	utils.ScanPorts(protocol, hostname, startPort, endPort, &wg)

	wg.Wait()
	log.Println("Scan complete.")
}
