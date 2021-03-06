package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Request with timeout

func main() {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Make request
	response, err := client.Get("https://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Copy data from the response to standard output
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of bytes copied to STDOUT:", n)
}
