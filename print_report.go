package main

import (
	"fmt"
	"slices"
)

type pageCount struct {
	page  string
	count int
}

func slicePages(pages map[string]int) (pageCounts []pageCount) {
	for k, v := range pages {
		pc := pageCount{
			page:  k,
			count: v,
		}

		pageCounts = append(pageCounts, pc)
	}

	slices.SortFunc(pageCounts, func(pc1, pc2 pageCount) int { return pc2.count - pc1.count })
	return
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`=============================
  REPORT for %v
=============================
`, baseURL)
	sliced := slicePages(pages)
	for _, p := range sliced {
		fmt.Printf("Found %v internal links to %v\n", p.count, p.page)
	}
}
