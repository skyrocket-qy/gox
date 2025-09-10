package utils

import (
	"context"
	"fmt"
	"log"
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

	dialer := &net.Dialer{
		Timeout: 1 * time.Second,
	}

	conn, err := dialer.DialContext(context.Background(), protocol, address)
	if err == nil {
		log.Printf("Port %d is open (%s)\n", port, protocol)

		if err := conn.Close(); err != nil {
			// Log the error, as we can't return it from a goroutine
			_ = err // Suppress the error if not logging
		}
	}
}
