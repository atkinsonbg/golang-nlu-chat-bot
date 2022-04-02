package classification

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/navossoc/bayesian"
)

const (
	Order   bayesian.Class = "Order"
	Hours   bayesian.Class = "Hours"
	Unknown bayesian.Class = "Unknown"
)

// Train will train a new classification model based on the JSON data in the "data/intents" folder
func Train(testPhrase []string) {
	// create a new classifier from the known intents
	classifier := bayesian.NewClassifierTfIdf(Order, Hours, Unknown)

	// load all the intent data from the JSON files
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

	// train the classifier on the loaded intents
	classifier.Learn(orderIntents, Order)
	classifier.Learn(hoursIntents, Hours)
	classifier.Learn(unknownIntents, Unknown)

	// convert terms freq to TFIDF
	classifier.ConvertTermsFreqToTfIdf()

	// save the trained classifier
	classifier.WriteToFile("models/classifier")

	// test the model on a supplied test sentence
	fmt.Printf("TEST SENTENCE: %s\n", testPhrase)

	cleanedTest := SentencesToWords(testPhrase)
	fmt.Printf("CLEANED SENTENCE: %s\n", cleanedTest)

	scores, likely, _ := classifier.LogScores(
		cleanedTest,
	)

	// print the results
	fmt.Printf("Scores: %.2f \n", scores)
	fmt.Printf("Likely Index: %d, Class: %s \n", likely, classifier.Classes[likely])
}

// GetIntents takes a supplied filename, JSON file, and reads its contents, running the sentences
// through the SentenceToWords util and returns a string array of words
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
