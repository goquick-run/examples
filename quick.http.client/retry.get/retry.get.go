package main

import (
	"fmt"
	"log"

	"github.com/jeffotoni/quick/http/client"
)

func main() {

	// Retry Delay Format Support
	//
	// The retry delay parameter supports various formats for flexibility in defining the wait time between retries.
	// Additionally, it allows enabling exponential backoff by appending "-bex".
	//
	// Supported formats:
	//
	// - "2mil"      => 2 milliseconds
	// - "2s"        => 2 seconds
	// - "2min"      => 2 minutes
	//
	// When using the "-bex" suffix, the delay will increase exponentially on each retry attempt
	// using the formula: waitTime = baseDelay * (2^attempt).
	//
	// Example Usage:
	//
	// client.WithRetry(
	// 		3,                 // Maximum number of retries
	// 		"2s",              // Delay between attempts
	// 		true,              // Use exponential backoff
	// 		"500,502,503,504", // HTTP status for retry
	// 		true,              // show Logger
	// 	),
	//
	// This configuration will retry up to 3 times with an exponential backoff starting at 2 seconds,
	// and will only retry if the response status is 500, 502, 503, or 504.
	cClient := client.New(
		client.WithRetry(
			3,                 // Maximum number of retries
			"2s",              // Delay between attempts
			true,              // Use exponential backoff
			"500,502,503,504", // HTTP status for retry
			true,              // show Logger
		),
	)

	resp, err := cClient.Get("http://localhost:3000/v1/user/1234")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET response:", string(resp.Body))
}
