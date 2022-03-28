package generators

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
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

	// Toppings file
	toppingsFile, err := os.Open(fmt.Sprintf("%s/data/entities/toppings.json", path))
	if err != nil {
		fmt.Println(err)
	}
	defer toppingsFile.Close()

	toppingsBytes, _ := ioutil.ReadAll(toppingsFile)
	var toppings []string
	err = json.Unmarshal(toppingsBytes, &toppings)
	if err != nil {
		fmt.Println(err)
	}

	// count determines how many toppings to include in a sentence
	count := 1

	sentences := []string{}
		for _, crust := range crusts {
			for _, size := range sizes {
				for _, order := range orders {
				order = strings.Replace(order, "{size}", size, -1)
				order = strings.Replace(order, "{crust}", crust, -1)

				// Figure out a random number based on the length of the toppings array
				max := len(toppings) - 1
				ran := rand.Intn(max-0) + 0

				topping := ""

				// if the random number + count is lower than the upper bounds of the array, fetch that many toppings
				if (ran + count) <= len(toppings) {
					tArr := toppings[ran:ran + count]
					topping = strings.Join(tArr[:], ",")
				} else {
					topping = toppings[ran]
				}

				order = strings.Replace(order, "{toppings}", topping, -1)

				sentences = append(sentences, order)
			}

			// up the count at the end of the 'sizes' loop as there are fewer sizes
			count++
		}

		// reset the count back to 1 once we leave the 'sizes' loop
		count = 1
	}

	output := new(strings.Builder)
	json.NewEncoder(output).Encode(sentences)
	fmt.Println(output.String())
}
