package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

const (
	version         = "1.0.0"
	defaultFileName = "credentials.json"
)

func main() {
	// Define the flag with default value and description
	fileName := flag.String("f", defaultFileName, "Path to the credentials file")
	outFileName := flag.String("o", "", "Path to save the base64 encoded output")
	showVersion := flag.Bool("v", false, "Show application version")

	// Parse the flags
	flag.Parse()

	// Show version and exit if version flag is set
	if *showVersion {
		fmt.Printf("%s\n", version)
		return
	}

	// Read the file content
	content, err := os.ReadFile(*fileName)
	if err != nil {
		fmt.Printf("Failed to read file: %s", err)
	}

	// Parse the content as JSON
	var jsonData interface{}
	if err = json.Unmarshal(content, &jsonData); err != nil {
		fmt.Printf("Failed to parse JSON: %s", err)
	}

	// Encode the JSON content to base64
	var jsonBytes []byte
	jsonBytes, err = json.Marshal(jsonData)
	if err != nil {
		fmt.Printf("Failed to marshal JSON: %s", err)
	}
	base64Content := base64.StdEncoding.EncodeToString(jsonBytes)

	// Save the base64 encoded content to a file if out flag is set
	if *outFileName != "" {
		if err = os.WriteFile(*outFileName, []byte(base64Content), 0644); err != nil {
			fmt.Printf("Failed to write to output file: %s", err)
		}
		fmt.Printf("Base64 encoded content saved to %s\n", *outFileName)
	} else {
		// Print the base64 encoded JSON content
		// fmt.Println("Base64 Encoded JSON content:")
		fmt.Println(base64Content)
	}
}
