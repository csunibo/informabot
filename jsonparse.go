package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type AutoReply []struct {
	Text  string `json:"text"`
	Reply string `json:"reply"`
}

func ParseAutoReplies() (AutoReply, error) {
	jsonFile, err := os.Open("./json/autoreply.json")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var autoreplies AutoReply
	json.Unmarshal(byteValue, &autoreplies)

	return autoreplies, nil
}

func ParseActions() error {
	jsonFile, err := os.Open("./json/actions.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var JSON map[string]interface{}
	json.Unmarshal(byteValue, &JSON)
	// TODO: finire il parsing delle azioni
	// fmt.Println(JSON)

	for key, value := range JSON {
		fmt.Println(key)
		fmt.Println(value)
		fmt.Printf("%T", value)
		break
	}
	// neededOutput := jsonToMap(JSON)
	// fmt.Println(neededOutput)

	return nil
}

// JSON TO MAP
// From https://stackoverflow.com/questions/70937676/extract-all-json-keys-and-values-as-map
func jsonToMap(data map[string]interface{}) map[string][]string {
	// final output
	out := make(map[string][]string)

	// check all keys in data
	for key, value := range data {
		// check if key not exist in out variable, add it
		if _, ok := out[key]; !ok {
			out[key] = []string{}
		}

		if valueA, ok := value.(map[string]interface{}); ok { // if value is map
			out[key] = append(out[key], "")
			for keyB, valueB := range jsonToMap(valueA) {
				if _, ok := out[keyB]; !ok {
					out[keyB] = []string{}
				}
				out[keyB] = append(out[keyB], valueB...)
			}
		} else if valueA, ok := value.([]interface{}); ok { // if value is array
			for _, valueB := range valueA {
				if valueC, ok := valueB.(map[string]interface{}); ok {
					for keyD, valueD := range jsonToMap(valueC) {
						if _, ok := out[keyD]; !ok {
							out[keyD] = []string{}
						}
						out[keyD] = append(out[keyD], valueD...)
					}
				} else {
					out[key] = append(out[key], fmt.Sprintf("%v", valueB))
				}
			}
		} else { // if string and numbers and other ...
			out[key] = append(out[key], fmt.Sprintf("%v", value))
		}
	}
	return out
}
