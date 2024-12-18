package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func main2() {
	start := time.Now()
	ch := make(chan string)
	var wg sync.WaitGroup

	for _, url := range os.Args[1:] {
		wg.Add(1) // Increment the WaitGroup counter
		go func(url string) {
			defer wg.Done() // Decrement the counter when done
			fetch(url, ch)
		}(url)
	}

	go func() {
		wg.Wait() // Wait for all goroutines to finish
		close(ch) // Close the channel to signal completion
	}()

	for msg := range ch {
		fmt.Println(msg)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
