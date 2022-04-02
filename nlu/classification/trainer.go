package classification

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/bbalet/stopwords"
	"github.com/navossoc/bayesian"
)

const (
	Order   bayesian.Class = "Order"
	Hours   bayesian.Class = "Hours"
	Unknown bayesian.Class = "Unknown"
)

func Train(testPhrase []string) {
	classifier := bayesian.NewClassifierTfIdf(Order, Hours, Unknown)

	hoursIntents, err := GetIntents("data/intents/hours.json")
	if err != nil {
		log.Println(err)
	}

	orderIntents, err := GetIntents("data/intents/order.json")
	if err != nil {
		log.Println(err)
	}

	unknownIntents, err := GetIntents("data/intents/unknown.json")
	if err != nil {
		log.Println(err)
	}

	classifier.Learn(orderIntents, Order)
	classifier.Learn(hoursIntents, Hours)
	classifier.Learn(unknownIntents, Unknown)

	classifier.ConvertTermsFreqToTfIdf()

	classifier.WriteToFile("models/classifier")

	fmt.Printf("TEST SENTENCE: %s\n", testPhrase)
	cleanedTest := SentencesToWords(testPhrase)
	fmt.Printf("CLEANED SENTENCE: %s\n", cleanedTest)
	scores, likely, _ := classifier.LogScores(
		cleanedTest,
	)

	fmt.Printf("Scores: %.2f \n", scores)
	fmt.Printf("Likely Index: %d, Class: %s \n", likely, classifier.Classes[likely])
}

func GetIntents(filename string) ([]string, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var intents []string
	err = json.Unmarshal(byteValue, &intents)
	if err != nil {
		return nil, err
	}

	words := SentencesToWords(intents)

	return words, nil
}

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

func LoadClassifier() (*bayesian.Classifier, error) {
	classifier, err := bayesian.NewClassifierFromFile("models/classifier")
	if err != nil {
		return nil, err
	}

	return classifier, nil
}
