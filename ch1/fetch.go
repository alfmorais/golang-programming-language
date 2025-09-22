package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, baseUrl := range os.Args[1:] {
		prefix := "https://"
		url := baseUrl

		if !strings.HasPrefix(baseUrl, prefix) {
			fmt.Print("baseUrl: ", baseUrl, "\n")
			url = prefix + baseUrl
			fmt.Print("url: ", url, "\n")
		}
		resp, err := http.Get(url)
		fmt.Print(url, "\n")

		if err != nil {
			fmt.Fprint(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		body, err := io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Fprint(os.Stderr, "Fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", body, "\n")
		fmt.Println("Status Code: ", resp.StatusCode)
	}
}
