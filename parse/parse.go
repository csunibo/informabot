package parse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/csunibo/informabot/model"
	"github.com/mitchellh/mapstructure"
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

func ParseActions() ([]model.Action, error) {
	jsonFile, err := os.Open("./json/actions.json")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return ParseActionsBytes(byteValue)
}

func ParseActionsBytes(bytes []byte) ([]model.Action, error) {
	var mapData map[string]interface{}
	json.Unmarshal(bytes, &mapData)

	var actions []model.Action
	for key, value := range mapData {
		action := model.GetActionFromType(value.(map[string]interface{})["type"].(string))
		action.Name = key
		err := mapstructure.Decode(value, &action)
		if err != nil {
			return nil, err
		}

		actions = append(actions, action)
	}

	return actions, nil
}
