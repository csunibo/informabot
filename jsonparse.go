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

func ParseActions() ([]Action, error) {
	jsonFile, err := os.Open("./json/actions.json")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var JSON map[string]interface{}
	json.Unmarshal(byteValue, &JSON)

	var actions []Action
	for key, value := range JSON {
		action := GetActionFromType(key)
		byteData, _ := json.Marshal(value)
		json.Unmarshal(byteData, &action)
		action.Name = key

		actions = append(actions, action)
	}

	return actions, nil
}
