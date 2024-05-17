package index

import (
	"gosearch/crawler/pkg/crawler"
	"strings"
)

type ReverseIndex struct {
	storage map[string][]int
}

func NewReverseIndex() *ReverseIndex {
	return &ReverseIndex{
		storage: make(map[string][]int),
	}
}

func (r *ReverseIndex) CreateIndex(docs *[]crawler.Document) map[string][]int {
	indexes := map[string][]int{}
	for _, doc := range *docs {
		tokens := strings.Split(doc.Title, " ")
		for _, token := range tokens {
			if token == "." || token == "-" || token == "&" {
				continue
			}

			_, ok := indexes[token]

			if !ok {
				indexes[token] = []int{doc.ID}
				continue
			}

			indexes[token] = append(indexes[token], doc.ID)
		}
	}

	return indexes
}
