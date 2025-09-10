package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Start the peer server to handle file-sharing requests.
func StartServer(port string) {
	lc := net.ListenConfig{}

	ln, err := lc.Listen(context.Background(), "tcp", "localhost:"+port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	defer func() {
		if cerr := ln.Close(); cerr != nil {
			log.Printf("Error closing listener: %v", cerr)
		}
	}()

	log.Printf("Server listening on port %s...\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)

			continue
		}

		go handleConnection(conn)
	}
}

// Handle incoming connection and file transfer request.
func handleConnection(conn net.Conn) {
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			log.Printf("Error closing connection: %v", cerr)
		}
	}()

	// Receive the file request (name of the file)
	var fileName string

	_, err := fmt.Fscan(conn, &fileName)
	if err != nil {
		log.Println("Error reading file name:", err)

		return
	}

	// Check if the requested file exists
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			if _, err := conn.Write([]byte("File not found")); err != nil {
				log.Printf("Error writing 'File not found' to connection: %v", err)
			}

			return
		}

		log.Println("Error checking file:", err)

		return
	}

	// Open the file and send it to the requesting peer
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error opening file:", err)

		return
	}

	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("Error closing file: %v", cerr)
		}
	}()

	// Send the file content to the client
	_, err = io.Copy(conn, file)
	if err != nil {
		log.Println("Error sending file:", err)
	}

	log.Printf("Sent file: %s\n", fileName)
}
