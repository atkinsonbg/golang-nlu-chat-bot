package main

import (
	"fmt"
	"github.com/navossoc/bayesian"
)

const (
	Good bayesian.Class = "Good"
	Bad bayesian.Class = "Bad"
	Weird bayesian.Class = "Weird"
)

func main() {
	// Create a classifier with TF-IDF support.
	classifier := bayesian.NewClassifierTfIdf(Good, Bad, Weird)

	goodStuff := []string{"tall", "rich", "handsome"}
	badStuff  := []string{"poor", "smelly", "ugly"}
	weirdStuff  := []string{"strange", "weird", "cthulu"}

	classifier.Learn(goodStuff, Good)
	classifier.Learn(badStuff,  Bad)
	classifier.Learn(weirdStuff,  Weird)

	// Required
	classifier.ConvertTermsFreqToTfIdf()

	scores, likely, _ := classifier.LogScores(
		[]string{"tall", "girl"},
	)

	fmt.Printf("Scores: %s", scores)
	fmt.Printf("Likely: %s", likely)
}
