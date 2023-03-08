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
