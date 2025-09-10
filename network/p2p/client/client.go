package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Connect to a peer and request a file.
func RequestFile(peerAddr, fileName string) (err error) {
	dialer := &net.Dialer{}

	conn, err := dialer.DialContext(context.Background(), "tcp", peerAddr)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := conn.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	// Send the file request to the peer
	if _, err := fmt.Fprintln(conn, fileName); err != nil {
		return err
	}

	fileContent, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer func() {
		if cerr := fileContent.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	_, err = io.Copy(fileContent, conn)
	if err != nil {
		return err
	}

	log.Printf("File downloaded: %s\n", fileName)

	return nil
}
