package main

import (
	"bufio"
	"fmt"
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
		fmt.Print("Please enter the Action: ")
		scanner.Scan()

		action = scanner.Text()
		if exp.MatchString(action) {
			fileName := action[4:]

			if fileName == "" {
				fmt.Println("Filename cannot be empty")
			} else {
				fmt.Printf("Requested filename: %s\n", fileName)
				client.RequestFile(peerAddr, fileName)
			}
		} else {
			fmt.Println("Invalid action. Please enter a valid action like 'get <filename>'")
		}
	}
}
