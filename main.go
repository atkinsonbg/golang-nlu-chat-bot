package main

import (
	"encoding/json"
	"fmt"
	"github.com/navossoc/bayesian"
	"io/ioutil"
	"log"
	"os"
)

const (
	Order bayesian.Class = "Order"
	Hours bayesian.Class = "Hours"
	Unknown bayesian.Class = "Unknown"
)

type Intent struct {
	Name	bayesian.Class
	Filename	string
}

func main() {
	//intents := []Intent{}
	//intents = append(intents, Intent{
	//	"Order",
	//	"data/intents/order.json",
	//}, Intent{
	//	"Hours",
	//	"data/intents/hours.json",
	//}, Intent{
	//	"Uknown",
	//	"data/intents/unknown.json",
	//})

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
	classifier.Learn(hoursIntents,  Hours)
	classifier.Learn(unknownIntents,  Unknown)

	classifier.ConvertTermsFreqToTfIdf()

	classifier.WriteToFile("classifier")

	scores, likely, _ := classifier.LogScores(
		[]string{"I need a pizza with regular crust pepperoni and extra cheese."},
	)

	fmt.Printf("Scores: %.2f \n", scores)
	fmt.Printf("Likely Index: %d \n", likely)
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

	return intents, nil
}

func LoadClassifier() (*bayesian.Classifier, error) {
	classifier, err := bayesian.NewClassifierFromFile("classifier")
	if err != nil {
		return nil, err
	}

	return classifier, nil
}
