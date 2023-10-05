package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	inputURL := "https://goo.gl/j7bP7u"

	err := validateAndResolveURL(inputURL)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func validateAndResolveURL(inputURL string) error {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return err
	}

	if !isValidURL(parsedURL) {
		return fmt.Errorf("Invalid URL: %s", inputURL)
	}

	originalURL, err := resolveURL(parsedURL)
	if err != nil {
		return err
	}

	fmt.Println("Original URL:", originalURL)
	return nil
}

func isValidURL(parsedURL *url.URL) bool {
	return parsedURL.Scheme == "http" || parsedURL.Scheme == "https"
}

func resolveURL(parsedURL *url.URL) (string, error) {
	resp, err := http.Head(parsedURL.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if len(resp.Header["Location"]) > 0 {
		return resp.Header["Location"][0], nil
	}

	return parsedURL.String(), nil
}
