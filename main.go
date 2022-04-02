package main

import (
	"flag"
	"fmt"
	"main/data/intents/generators"
	"main/nlu/classification"
)

func main() {
	trainFlag := flag.Bool("train", false, "Train a new model.")
	genFlag := flag.Bool("generate", false, "Generate data for training.")
	sentenceFlag := flag.String("sentence", "Blah", "Test message to use after training.")

	classifyFlag := flag.Bool("classify", false, "Classify a sentence using a prebuilt classifier.")

	flag.Parse()

	// Train a new model
	if *trainFlag {
		t := []string{*sentenceFlag}
		classification.Train(t)
	}

	// Generate sentences
	if *genFlag {
		generators.GenerateOrderSentences()
	}

	// Classify a sentence
	if *classifyFlag {
		res := classification.ClassifySentence(*sentenceFlag)
		fmt.Printf("RESULT FROM CLASSIFICATION: %s\n", res)
	}
}
