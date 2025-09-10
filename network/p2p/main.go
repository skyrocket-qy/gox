package main

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"github.com/skyrocket-qy/gox/network/p2p/client"
	"github.com/skyrocket-qy/gox/network/p2p/server"
)

// Struct to represent a peer.
type Peer struct {
	Addr string
	Port string
}

func main() {
	// Example peer server and peer client setup
	peerAddr := "localhost:8000" // Address of another peer to request files from
	serverPort := "8001"         // Port to listen on for incoming requests

	go server.StartServer(serverPort)

	exp := regexp.MustCompile("^get .+")

	var action string

	scanner := bufio.NewScanner(os.Stdin)

	for {
		log.Print("Please enter the Action: ")
		scanner.Scan()

		action = scanner.Text()
		if exp.MatchString(action) {
			fileName := action[4:]

			if fileName == "" {
				log.Println("Filename cannot be empty")
			} else {
				log.Printf("Requested filename: %s\n", fileName)

				if err := client.RequestFile(peerAddr, fileName); err != nil {
					log.Printf("Error requesting file: %v", err)
				}
			}
		} else {
			log.Println("Invalid action. Please enter a valid action like 'get <filename>'")
		}
	}
}
