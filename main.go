package main

import (
	"flag"
	"main/data/intents/generators"
	"main/nlu"
)

func main() {
	trainFlag := flag.Bool("train", false, "Train a new model.")
	genFlag := flag.Bool("generate", false, "Generate data for training.")
	flag.Parse()

	// Train a new model
	if *trainFlag {
		nlu.Train()
	}

	// Generate sentences
	if *genFlag {
		generators.GenerateOrderSentences()
	}
}