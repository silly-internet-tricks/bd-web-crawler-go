package main

import (
	"fmt"
	"log"
	"os"
)

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
	p, err := crawlPage(os.Args[1], os.Args[1], pages)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	for key, value := range p {
		fmt.Printf("I found *%v* links to the page: %v\n", value, key)
	}
}
