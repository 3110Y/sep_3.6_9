package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const name = "data-20190514T0100.json"

type Item struct {
	GlobalId uint64 `json:"global_id"`
}

func ParseFileToJson(fileName string) ([]Item, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return []Item{}, err
	}
	defer file.Close()
	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	err = json.Unmarshal(jsonData, &items)
	if err != nil {
		return items, err
	}
	return items, nil
}

func main() {
	items, err := ParseFileToJson(name)
	if err != nil {
		panic(fmt.Sprintf("Unable to read: %v\n", err))
	}
	var sum uint64
	for _, item := range items {
		sum += item.GlobalId
	}
	fmt.Println(sum)

}
