package main

import (
	"flag"
	"fmt"
	"gosearch/crawler1/pkg/crawler"
	"gosearch/crawler1/pkg/crawler/spider"
	"strings"
)

const depth int = 2

var needle string

func flagInit() {
	flag.StringVar(&needle, "needle", "", "Url search key")
	flag.Parse()
}

func main() {
	flagInit()
	if needle == "" {
		fmt.Println("Needle flag is empty")
		return
	}

	scanner := spider.New()
	resources := [...]string{"https://golang-org.appspot.com/", "https://go.dev/"}
	var scanResult []crawler.Document

	for _, doc := range resources {
		document, err := scanner.Scan(doc, depth)
		if err != nil {
			fmt.Printf("Scan not resolved: %s", err)
		}

		scanResult = append(scanResult, document...)
	}

	links := searchUrl(&needle, &scanResult)
	if len(links) == 0 {
		fmt.Println("Scan result is empty")
		return
	}

	for _, link := range links {
		fmt.Println(link)
	}
}

func searchUrl(needle *string, scanResult *[]crawler.Document) []string {
	var links []string

	for _, doc := range *scanResult {
		if strings.Contains(doc.URL, *needle) {
			links = append(links, doc.URL)
		}
	}

	return links
}
