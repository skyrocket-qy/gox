package main

import (
	"fmt"
	"sync"

	"github.com/skyrocket-qy/gox/Go/network/network-scanner/utils"
)

func main() {
	var hostname, protocol string
	var startPort, endPort int

	fmt.Print("Enter hostname or IP: ")
	fmt.Scanln(&hostname)
	fmt.Print("Enter protocol (tcp or tcp6): ")
	fmt.Scanln(&protocol)
	fmt.Print("Enter start port: ")
	fmt.Scanln(&startPort)
	fmt.Print("Enter end port: ")
	fmt.Scanln(&endPort)

	fmt.Printf("Scanning %s (%s) for open ports (%d-%d)...\n", hostname, protocol, startPort, endPort)

	var wg sync.WaitGroup

	utils.ScanPorts(protocol, hostname, startPort, endPort, &wg)

	wg.Wait()
	fmt.Println("Scan complete.")
}
