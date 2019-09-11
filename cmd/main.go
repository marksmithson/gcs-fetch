/*
Fetch a file from Google Cloud Storage

Usage:
	gcs-fetch gs://bucket/object output-file
 */
package main

import (
	"fmt"
	"github.com/marksmithson/cloudstorage-fetch/internal/pkg/gcsfetch"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		printUsage()
		os.Exit(1)
	}

	gsObject := os.Args[1]
	outputFilename := os.Args[2]

	if gsObject == "" || outputFilename == "" {
		printUsage()
		os.Exit(1)
	}

	bytesRead, err := gcsfetch.Fetch(gsObject, outputFilename)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Copied %d Bytes\n", bytesRead)
}

func printUsage() {
	fmt.Println("Usage: gcs-fetch gs://bucket/object output-file")
}