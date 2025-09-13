package netscan

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanPorts(t *testing.T) {
	// Redirect log output to a buffer
	var logBuf bytes.Buffer
	log.SetOutput(&logBuf)

	t.Run("TCP4", func(t *testing.T) {
		logBuf.Reset()
		// Start a local TCP listener on a random port
		listener, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			t.Fatalf("Failed to listen on a port: %v", err)
		}
		defer listener.Close()
		openPort := listener.Addr().(*net.TCPAddr).Port

		var wg sync.WaitGroup
		ScanPorts("tcp", "127.0.0.1", openPort, openPort, &wg)
		wg.Wait()

		// Check if the open port was logged
		assert.Contains(t, logBuf.String(), fmt.Sprintf("Port %d is open (tcp)", openPort))
	})

	t.Run("TCP6", func(t *testing.T) {
		logBuf.Reset()
		// Start a local TCP6 listener on a random port
		listener6, err := net.Listen("tcp6", "[::1]:0")
		if err != nil {
			// Skip if IPv6 is not supported
			if strings.Contains(err.Error(), "address family not supported by protocol") {
				t.Skip("IPv6 not supported, skipping test")
			}
			t.Fatalf("Failed to listen on a port: %v", err)
		}
		defer listener6.Close()
		openPort6 := listener6.Addr().(*net.TCPAddr).Port

		var wg6 sync.WaitGroup
		ScanPorts("tcp6", "localhost", openPort6, openPort6, &wg6)
		wg6.Wait()

		// Check if the open port was logged
		assert.Contains(t, logBuf.String(), fmt.Sprintf("Port %d is open (tcp6)", openPort6))
	})

	t.Run("Closed Port", func(t *testing.T) {
		logBuf.Reset()
		var wgClosed sync.WaitGroup
		// Assuming port 1 is not open
		ScanPorts("tcp", "127.0.0.1", 1, 1, &wgClosed)
		wgClosed.Wait()
		assert.Empty(t, logBuf.String(), "Expected no log output for a closed port")
	})
}
