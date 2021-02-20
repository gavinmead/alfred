package tags

import (
	"testing"

	"github.com/aaaton/golem/v4"
	"github.com/aaaton/golem/v4/dicts/en"
	"github.com/jdkato/prose/v2"
)

func Test_TagExtract(t *testing.T) {
	doc, err := prose.NewDocument("@alfred create the following tags: test1, test2 and test3\n")
	if err != nil {
		t.Fatal(err)
	}

	//Iterate over the doc's tokens:
	for _, tok := range doc.Tokens() {
		t.Log(tok.Text, tok.Tag, tok.Label)
		// Go NNP B-GPE
		// is VBZ O
		// an DT O
		// ...
	}

	// Iterate over the doc's named-entities:
	// for _, ent := range doc.Entities() {
	// 	t.Log(ent.Text, ent.Label)
	// 	// Go GPE
	// 	// Google GPE
	// }

	// Iterate over the doc's sentences:
	for _, sent := range doc.Sentences() {
		t.Log(sent.Text)
		// Go is an open-source programming language created at Google.
	}

	lemmatizer, err := golem.New(en.New())
	if err != nil {
		panic(err)
	}
	word := lemmatizer.Lemma("tags")
	t.Log(word)
}
