package index

type ReverseIndex struct {
	storage map[string][]int
}

func NewReverseIndex() *ReverseIndex {
	return &ReverseIndex{
		storage: make(map[string][]int),
	}
}

//func (r *ReverseIndex) CreateIndex(docs *[]crawler.Document) map[string][]int {
//	index := map[string][]int{}
//	var tokens []string
//	for {
//
//	}
//}
