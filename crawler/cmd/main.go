package main

import (
	"flag"
	"fmt"
	"gosearch/crawler/pkg/crawler"
	"gosearch/crawler/pkg/crawler/spider"
	"gosearch/crawler/pkg/index"
	"math/rand"
	"sort"
	"time"
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

	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	for i := range scanResult {
		scanResult[i].ID = randomizer.Int()
	}

	sort.Slice(scanResult, func(i, j int) bool {
		return scanResult[i].ID < scanResult[j].ID
	})

	indexBuilder := index.NewReverseIndex()
	indexes := indexBuilder.CreateIndex(&scanResult)

	linkIds, ok := indexes[needle]
	if !ok {
		fmt.Printf("Link with work %s not found", needle)
		return
	}

	for _, id := range linkIds {
		link := sort.Search(len(scanResult), func(i int) bool {
			return scanResult[i].ID >= id
		})
		if link < len(scanResult) && scanResult[link].ID == id {
			fmt.Println(scanResult[link].URL)
		}
	}
}
