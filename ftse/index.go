package ftse

import (
	"encoding/json"
	"fmt"
	"os"
)

// Index is the implementation of a search index. (inverted index, map tokens to document IDs)
type Index map[string][]int

// New creates a new search index.
func NewIndex() Index {
	return make(Index)
}

// Add adds documents to the index.
func (idx Index) Add(docs []Document) {
	for _, doc := range docs {
		for _, token := range Analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				// Don't add same ID twice.
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

// Generate a search index file.
func GenIndex(docs []Document) {
	idx := make(Index)
	for _, doc := range docs {
		for _, token := range Analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				// Don't add same ID twice.
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
	store, err := json.Marshal(idx)
    if err != nil {
        fmt.Println("Transfer to json failed:", err)
        return
    }
	os.WriteFile("index.json", store, 0777)
}

// intersection returns the set intersection between a and b.
// a and b have to be sorted in ascending order and contain no duplicates.
func intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

// Search queries the index for the given text.
func (idx Index) Search(text string) []int {
	var r []int
	for _, token := range Analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = intersection(r, ids)
			}
		} else {
			// Token doesn't exist.
			return nil
		}
	}
	return r
}

// Qsearch queries the index for the given text.
func Qsearch(idx Index, docs []Document, text string) int{
	var matchedIDs []int
	for _, token := range Analyze(text) {
		if ids, ok := idx[token]; ok {
			if matchedIDs == nil {
				matchedIDs = ids
			} else {
				matchedIDs = intersection(matchedIDs, ids)
			}
		}
	}

	for _, id := range matchedIDs {
		doc := docs[id]
		fmt.Printf("%d\t%s\n", id, doc.Text)
	}
	
	return len(matchedIDs)
}