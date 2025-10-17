# coder/gzip

The `coder/gzip` package provides a concrete implementation of the `coder.Coder` interface, offering GZIP compression and decompression functionalities. It utilizes Go's standard `compress/gzip` package to efficiently compress and decompress byte data.

## Features

*   **GZIP Compression:** Compresses byte arrays using the GZIP algorithm.
*   **GZIP Decompression:** Decompresses GZIP-encoded byte arrays back to their original form.
*   **`coder.Coder` Interface Implementation:** Adheres to the generic `coder.Coder` interface, enabling interchangeable use with other `Coder` implementations.

## Usage Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/skyrocket-qy/ciri/coder"
	"github.com/skyrocket-qy/ciri/coder/gzip"
)

func main() {
	// Create a new GZIP Coder
	var gzipCoder coder.Coder = gzip.New()

	originalData := []byte("This is a sample string that will be compressed using the GZIP algorithm. It should become smaller after compression.")
	fmt.Printf("Original Data Length: %d bytes\n", len(originalData))
	fmt.Printf("Original Data: %s\n\n", originalData)

	// Encode (compress) the data
	encodedData, err := gzipCoder.Encoder(originalData)
	if err != nil {
		log.Fatalf("Error encoding data: %v", err)
	}
	fmt.Printf("Encoded Data Length: %d bytes\n", len(encodedData))
	fmt.Printf("Encoded Data (first 50 bytes): %x...\n\n", encodedData[:50])

	// Decode (decompress) the data
	decodedData, err := gzipCoder.Decoder(encodedData)
	if err != nil {
		log.Fatalf("Error decoding data: %v", err)
	}
	fmt.Printf("Decoded Data Length: %d bytes\n", len(decodedData))
	fmt.Printf("Decoded Data: %s\n\n", decodedData)

	// Verify if the decoded data matches the original
	if string(originalData) == string(decodedData) {
		fmt.Println("Verification successful: Decoded data matches original data.")
	} else {
		fmt.Println("Verification failed: Decoded data does NOT match original data.")
	}
}
```