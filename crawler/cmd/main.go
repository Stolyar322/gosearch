package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"gosearch/crawler/pkg/crawler"
	"gosearch/crawler/pkg/crawler/spider"
	"io"
	"os"
	"strings"
)

const depth int = 2

var needle string

func flagInit() {
	flag.StringVar(&needle, "needle", "", "Url search key")
	flag.Parse()
}

const filePath = "./../links.txt"

func main() {
	flagInit()
	if needle == "" {
		fmt.Println("Needle flag is empty")
		return
	}

	_, err := os.Stat(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fileStorageInit()
		} else {
			panic(err)
		}
	}

	links, err := search(filePath, needle)
	if err != nil {
		panic(err)
	}

	if len(links) == 0 {
		fmt.Printf("Links with %s not found", needle)
		return
	}

	for _, link := range links {
		fmt.Println(link)
	}
}

func search(pathFile string, needle string) (res []string, err error) {
	file, err := os.Open(pathFile)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}

	}(file)

	var links []string
	scanner := bufio.NewReader(file)
	for {
		buff, _, err := scanner.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return []string{}, err
		}

		link := fmt.Sprintf("%s", buff)
		if strings.Contains(link, needle) {
			links = append(links, link)
		}
	}

	return links, nil
}

func fileStorageInit() {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

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

	for _, result := range scanResult {
		_, err := file.Write([]byte(result.URL + "\n"))
		if err != nil {
			panic(err)
		}
	}
}
