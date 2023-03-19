package parse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/csunibo/informabot/model"
	"github.com/mitchellh/mapstructure"
)

func ParseAutoReplies() ([]model.AutoReply, error) {
	jsonFile, err := os.Open("./json/autoreply.json")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var autoreplies []model.AutoReply
	json.Unmarshal(byteValue, &autoreplies)

	return autoreplies, nil
}

func ParseSettings() (model.Settings, error) {
	jsonFile, err := os.Open("./json/settings.json")
	if err != nil {
		fmt.Println(err)
		return model.Settings{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var settings model.Settings
	json.Unmarshal(byteValue, &settings)

	return settings, nil
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
		action := model.GetActionFromType(key, value.(map[string]interface{})["type"].(string))
		err := mapstructure.Decode(value, &action)
		if err != nil {
			return nil, err
		}

		actions = append(actions, action)
	}

	return actions, nil
}

func ParseMemeList() ([]model.Meme, error) {
	jsonFile, err := os.Open("./json/memes.json")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var mapData map[string]interface{}
	json.Unmarshal(byteValue, &mapData)

	var memes []model.Meme
	for key, value := range mapData {
		meme := model.Meme{
			Name: key,
			Text: value.(string),
		}
		memes = append(memes, meme)
	}

	return memes, nil
}
