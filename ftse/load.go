package ftse

import (
	"encoding/json"
	"encoding/xml"
	
	"os"
)

// Document represents a Wikipedia article.
type Document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

// LoadDoc loads a Wikipedia article from a dump file.
func LoadDoc(filename string) ([]Document, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := xml.NewDecoder(f)
	dump := struct {
		Documents []Document `xml:"doc"`
	}{}
	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}

	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}

// LoadIndex loads index json file.
func LoadIndex(filename string) (Index, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	info := make(Index)
	if err := dec.Decode(&info); err != nil {
		return nil, err
	}

	return info, nil
}