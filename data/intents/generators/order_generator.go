package generators

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func GenerateOrderSentences() {
	// Current directory
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	// Order generator template file
	ordersFile, err := os.Open(fmt.Sprintf("%s/data/intents/generators/order_generator.json", path))
	if err != nil {
		fmt.Println(err)
	}
	defer ordersFile.Close()

	ordersBytes, _ := ioutil.ReadAll(ordersFile)
	var orders []string
	err = json.Unmarshal(ordersBytes, &orders)
	if err != nil {
		fmt.Println(err)
	}

	// Crusts file
	crustsFile, err := os.Open(fmt.Sprintf("%s/data/entities/crusts.json", path))
	if err != nil {
		fmt.Println(err)
	}
	defer crustsFile.Close()

	crustsBytes, _ := ioutil.ReadAll(crustsFile)
	var crusts []string
	err = json.Unmarshal(crustsBytes, &crusts)
	if err != nil {
		fmt.Println(err)
	}

	// Sizes file
	sizesFile, err := os.Open(fmt.Sprintf("%s/data/entities/sizes.json", path))
	if err != nil {
		fmt.Println(err)
	}
	defer sizesFile.Close()

	sizesBytes, _ := ioutil.ReadAll(sizesFile)
	var sizes []string
	err = json.Unmarshal(sizesBytes, &sizes)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(crusts)
	fmt.Println(sizes)
}
