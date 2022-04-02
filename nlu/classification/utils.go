package classification

import (
	"github.com/bbalet/stopwords"
	"log"
	"regexp"
	"strings"
)

// SentencesToWords converts a string array of sentences to a string array of words,
// it also performs a removal of stopwords from the sentences before splitting them.
func SentencesToWords(sentences []string) []string {
	w := []string{}

	for _, s := range sentences {
		// remove all non-alphanumeric characters
		reg, err := regexp.Compile("[^a-zA-Z0-9]+")
		if err != nil {
			log.Fatal(err)
		}
		s = reg.ReplaceAllString(s, " ")

		// remove any stopwords from the sentence
		s = stopwords.CleanString(s, "en", true)

		// lowercase all the text
		s = strings.ToLower(s)

		// split the sentence into words
		words := strings.Split(s, " ")
		for _, v := range words {
			// only add words with a length greater than 2
			if len(v) > 2 {
				w = append(w, v)
			}
		}
	}

	return w
}
