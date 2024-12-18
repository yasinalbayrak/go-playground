package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		httpPrefix := "http://"

		if strings.HasPrefix(url, httpPrefix) {
			url = httpPrefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		testCopy(resp)
		//testReadAll(resp)
	}
}

func testReadAll(resp *http.Response) {
	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", resp.Request.URL, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}

func testCopy(resp *http.Response) {
	_, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", resp.Request.URL, err)
	}

}
