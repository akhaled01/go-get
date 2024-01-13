package funcs

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

// this file contains all there is to it about mirroring
// we tokenize html, and recursively follow any links in it

func FetchHTML(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	htmlContent, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(htmlContent), nil
}

func ExtractResources(htmlContent, baseURL string) ([]string, error) {
	var resources []string

	//* tokenize html, mfairs style :)
	tokenizer := html.NewTokenizer(strings.NewReader(htmlContent))
	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return resources, nil
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "link" || token.Data == "script" || token.Data == "img" {
				for _, attr := range token.Attr {
					if attr.Key == "href" || attr.Key == "src" {
						resourceURL := attr.Val
						if !strings.HasPrefix(resourceURL, "http") {
							// Handle relative URLs
							resourceURL = baseURL + "/" + resourceURL
						}
						resources = append(resources, resourceURL)
					}
				}
			}
		}
	}
}

func DownloadResources(resources []string, baseURL string) error {
	for _, resourceURL := range resources {
		response, err := http.Get(resourceURL)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		// Extract hostname from the base URL
		parsedURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		hostname := parsedURL.Hostname()

		// Extract filename from URL
		tokens := strings.Split(resourceURL, "/")
		filename := tokens[len(tokens)-1]

		// Create directories if they don't exist
		directory := filepath.Join(hostname, strings.Join(tokens[2:len(tokens)-1], "/"))
		os.MkdirAll(directory, os.ModePerm)

		// Create and write the file
		filePath := filepath.Join(directory, filename)
		file, err := os.Create(filePath+"1")
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, response.Body)
		if err != nil {
			return err
		}

		fmt.Println("Downloaded:", resourceURL, " -> Saved to:", filePath)
	}
	return nil
}
