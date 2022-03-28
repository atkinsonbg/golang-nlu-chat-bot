package nlu

import (
	"encoding/json"
	"fmt"
	"github.com/navossoc/bayesian"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	Order bayesian.Class = "Order"
	Hours bayesian.Class = "Hours"
	Unknown bayesian.Class = "Unknown"
)

func Train() {
	classifier := bayesian.NewClassifierTfIdf(Order, Hours, Unknown)

	hoursIntents, err := GetIntents("data/intents/hours.json")
	fmt.Printf("Number of hours intents: %d \n", len(hoursIntents))
	if err != nil {
		log.Println(err)
	}

	orderIntents, err := GetIntents("data/intents/order.json")
	fmt.Printf("Number of orders intents: %d \n", len(orderIntents))
	if err != nil {
		log.Println(err)
	}

	unknownIntents, err := GetIntents("data/intents/unknown.json")
	fmt.Printf("Number of unknown intents: %d \n", len(unknownIntents))
	if err != nil {
		log.Println(err)
	}

	classifier.Learn(orderIntents, Order)
	classifier.Learn(hoursIntents,  Hours)
	classifier.Learn(unknownIntents,  Unknown)

	classifier.ConvertTermsFreqToTfIdf()

	//classifier.WriteToFile("classifier")

	test := []string{"Can I get a thin crust pizza with the following toppings garlic, size small?"}
	scores, likely, _ := classifier.LogScores(
		SentencesToWords(test),
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

func SentencesToWords(sentences []string) []string {
	w := []string{}

	for _, s := range sentences {
		words := strings.Split(s, " ")
		for _, v := range words {
			w = append(w, v)
		}
	}

	return w
}

func LoadClassifier() (*bayesian.Classifier, error) {
	classifier, err := bayesian.NewClassifierFromFile("classifier")
	if err != nil {
		return nil, err
	}

	return classifier, nil
}

