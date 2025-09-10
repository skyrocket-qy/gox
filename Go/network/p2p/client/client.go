package client

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Connect to a peer and request a file
func RequestFile(peerAddr, fileName string) error {
	conn, err := net.Dial("tcp", peerAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Send the file request to the peer
	fmt.Fprintln(conn, fileName)

	fileContent, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer fileContent.Close()

	_, err = io.Copy(fileContent, conn)
	if err != nil {
		return err
	}

	log.Printf("File downloaded: %s\n", fileName)
	return nil
}
