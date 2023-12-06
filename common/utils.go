package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const SessionCookieEnvKey = "AOC_SESSION_COOKIE"

func FetchAOCInput(day int) string {
	url := fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)
	session := os.Getenv(SessionCookieEnvKey)
	if session == "" {
		panic("session cookie not set in environment variables: AOC_SESSION_COOKIE")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// Set the session cookie
	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("failed to fetch data: status code %d", resp.StatusCode))
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("error reading response body: %v", err))
	}

	return string(body)
}
