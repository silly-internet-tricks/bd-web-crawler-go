package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"
)

const maxConcurrency = 10

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		log.Fatal("no website provided")
	}

	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		log.Fatal("too many arguments provided")
	}

	fmt.Printf("starting crawl of: %v\n", os.Args[1])
	pages := make(map[string]int)
	baseURL, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatalf("could not url parse the argument that was passed in!")
	}

	mu := &sync.Mutex{}
	concurrencyControl := make(chan struct{}, maxConcurrency)
	wg := &sync.WaitGroup{}
	cfg := config{
		pages,
		baseURL,
		mu,
		concurrencyControl,
		wg,
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(os.Args[1])
	cfg.wg.Wait()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	for key, value := range cfg.pages {
		fmt.Printf("I found *%v* links to the page: %v\n", value, key)
	}
}
