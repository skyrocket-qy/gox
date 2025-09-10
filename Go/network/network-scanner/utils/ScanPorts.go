package utils

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func ScanPorts(protocol, hostname string, startPort, endPort int, wg *sync.WaitGroup) {
	for port := startPort; port <= endPort; port++ {
		wg.Add(1)

		go scanPort(protocol, hostname, port, wg)
	}
}

func scanPort(protocol, hostname string, port int, wg *sync.WaitGroup) {
	defer wg.Done()

	var address string

	if protocol == "tcp6" {
		if hostname == "localhost" {
			hostname = "::1"
		}

		address = fmt.Sprintf("[%s]:%d", hostname, port)
	} else {
		address = fmt.Sprintf("%s:%d", hostname, port)
	}

	conn, err := net.DialTimeout(protocol, address, 1*time.Second)
	if err == nil {
		fmt.Printf("Port %d is open (%s)\n", port, protocol)
		conn.Close()
	}
}
