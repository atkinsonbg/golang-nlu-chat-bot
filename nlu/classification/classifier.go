package classification

import (
	"fmt"

	"github.com/navossoc/bayesian"
)

// ClassifySentence loads a pre-trained model and classifies the supplied sentence
func ClassifySentence(sentence string) string {
	// load the pre-trained model
	classifier, err := LoadClassifier()
	if err != nil {
		fmt.Printf("ERROR OCCURED LOADING CLASSIFIER: %s", err)
	}

	// split the sentence into cleaned words
	cleanedSentence := SentencesToWords([]string{sentence})

	// perform classification on the words
	scores, likely, _ := classifier.LogScores(
		cleanedSentence,
	)

	// print the results of classification
	fmt.Printf("Scores: %.2f \n", scores)
	fmt.Printf("Likely Index: %d, Class: %s \n", likely, classifier.Classes[likely])

	return string(classifier.Classes[likely])
}

// LoadClassifier loads a pre-trained classification model in the "models" folder
func LoadClassifier() (*bayesian.Classifier, error) {
	// load the pre-trained model from a file
	classifier, err := bayesian.NewClassifierFromFile("models/classifier")
	if err != nil {
		return nil, err
	}

	return classifier, nil
}
