package classification

import (
	"fmt"
	"github.com/navossoc/bayesian"
)

func ClassifySentence(sentence string) string {
	classifier, err := LoadClassifier()
	if err != nil {
		fmt.Printf("ERROR OCCURED LOADING CLASSIFIER: %s", err)
	}

	cleanedSentence := SentencesToWords([]string{sentence})

	scores, likely, _ := classifier.LogScores(
		cleanedSentence,
	)

	fmt.Printf("Scores: %.2f \n", scores)
	fmt.Printf("Likely Index: %d, Class: %s \n", likely, classifier.Classes[likely])

	return string(classifier.Classes[likely])
}

func LoadClassifier() (*bayesian.Classifier, error) {
	classifier, err := bayesian.NewClassifierFromFile("models/classifier")
	if err != nil {
		return nil, err
	}

	return classifier, nil
}
