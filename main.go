package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
)

const defaultMaxConcurrency = 10
const defaultMaxPages = 256

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		log.Fatal("no website provided")
	}

	if len(os.Args) > 4 {
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
	var maxConcurrency int
	var maxPages int
	if len(os.Args) > 2 {
		maxConcurrency, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("couldn't parse argument two as integer. using default max concurrency")
			maxConcurrency = defaultMaxConcurrency
		}
	} else {
		maxConcurrency = defaultMaxConcurrency
	}

	if len(os.Args) > 3 {
		maxPages, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("couldn't parse argument three as integer. using default max pages")
		}
	} else {
		maxPages = defaultMaxPages
	}

	concurrencyControl := make(chan struct{}, maxConcurrency)
	wg := &sync.WaitGroup{}
	cfg := config{
		maxPages,
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

	printReport(cfg.pages, cfg.baseURL.String())
}
