package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// ORDERS FILE
	ordersFile, err := os.Open("order_generator.json")
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

	// CRUSTS FILE
	crustsFile, err := os.Open("../../entities/crusts.json")
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

	fmt.Println(crusts)
}
