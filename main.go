package main

import (
	"flag"
	"main/data/intents/generators"
	"main/nlu"
)

func main() {
	trainFlag := flag.Bool("train", false, "Train a new model.")
	genFlag := flag.Bool("generate", false, "Generate data for training.")
	testPhraseFlag := flag.String("test", "Blah", "Test message to use after training.")

	flag.Parse()

	// Train a new model
	if *trainFlag {
		t := []string{*testPhraseFlag}
		nlu.Train(t)
	}

	// Generate sentences
	if *genFlag {
		generators.GenerateOrderSentences()
	}
}
