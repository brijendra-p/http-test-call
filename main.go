package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Define CLI flags
	url := flag.String("url", "", "Target API URL (required)")
	jsonData := flag.String("data", "", "JSON payload (required)")
	n := flag.Int("n", 0, "Number of requests (required)")
	waitTime := flag.Int("wait", 0, "Wait time in seconds between requests (required)")

	flag.Parse()

	// Validate required arguments
	if *url == "" || *jsonData == "" || *n <= 0 || *waitTime <= 0 {
		log.Fatal("Missing required arguments.")
		return
	}

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	logger := log.New(file, "", log.Lshortfile)

	logger.Printf("Sending requests to %s for %d times on interval of %d seconds with request data: %s", *url, *n, *waitTime, *jsonData)

	for i := range *n {
		logger.Printf("Sending request %d", i+1)

		resp, err := http.Post(*url, "application/json", bytes.NewBuffer([]byte(*jsonData)))
		if err != nil {
			logger.Printf("Error creating request %d: %v", i+1, err.Error())
			continue
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Printf("Error reading response %d: %v", i+1, err.Error())
			return
		}
		logger.Printf("response of request %d: %v", i+1, map[string]any{
			"status": resp.Status,
			"body":   string(body),
		})

		time.Sleep(time.Duration(*waitTime) * time.Second)
	}
}
