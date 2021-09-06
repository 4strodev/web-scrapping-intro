package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Request with timeout

//func main() {
//// Create HTTP client with timeout
//client := &http.Client{
//Timeout: 30 * time.Second,
//}

//// Make request
//response, err := client.Get("https://www.google.com/")
//if err != nil {
//log.Fatal(err)
//}
//defer response.Body.Close()

//// Copy data from the response to standard output
//n, err := io.Copy(os.Stdout, response.Body)
//if err != nil {
//log.Fatal(err)
//}

//fmt.Println("Number of bytes copied to STDOUT:", n)
//}

func main() {
	// Create HTTP client
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	// Create a custom request before sending
	request, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("User-Agent", "Not Firefox")

	// Make request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Copy data from the response to stdout
	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}
