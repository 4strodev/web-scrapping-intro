package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Make request
	response, err := http.Get("https://www.devdungeon.com/archive")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create output file
	outFile, err := os.Create("output.html")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	// Copy data from HTTP response to file
	_, err = io.Copy(outFile, response.Body)
	fmt.Printf("%v, %T", response.Body, response.Body)

	if err != nil {
		log.Fatal(err)
	}
}
