package stevens

import (
	"encoding/xml"
	"fmt"

	"github.com/robherley/stevens-course-api/util"
)

// TermList simple struct for terms
type TermList struct {
	Terms []Term `xml:"Term"`
}

// Term struct for a term from xml
type Term struct {
	Code string `xml:"Code,attr"`
	Name string `xml:"Name,attr"`
}

// FetchTerms fetches a list of terms
func FetchTerms() TermList {
	body, err := util.ByteRequest("https://web.stevens.edu/scheduler/core/core.php?cmd=terms")
	if err != nil {
		panic(err)
	}
	var ts TermList
	err = xml.Unmarshal(body, &ts)
	if err != nil {
		panic(err)
	}
	return ts
}

// ListTerms lists all the valid terms w/ out desciptions
func (ts TermList) ListTerms() []string {
	ls := make([]string, len(ts.Terms))
	for i, t := range ts.Terms {
		ls[i] = t.Code
	}
	return ls
}

func (ts TermList) print() {
	for _, t := range ts.Terms {
		fmt.Printf("[%s]: %s\n", t.Code, t.Name)
	}
}
