package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mitchellh/mapstructure"
	"golang.org/x/exp/slices"

	"github.com/csunibo/informabot/utils"
)

const groupsPath = "./json/groups.json"

func ParseAutoReplies() (autoReplies []AutoReply, err error) {
	file, err := os.Open("./json/autoreply.json")
	if err != nil {
		return nil, fmt.Errorf("error reading autoreply.json file: %w", err)
	}

	err = json.NewDecoder(file).Decode(&autoReplies)
	if err != nil {
		return nil, fmt.Errorf("error parsing autoreply.json file: %w", err)
	}

	return
}

func ParseSettings() (settings SettingsStruct, err error) {
	file, err := os.Open("./json/settings.json")
	if err != nil {
		return SettingsStruct{}, fmt.Errorf("error reading settings.json file: %w", err)
	}

	err = json.NewDecoder(file).Decode(&settings)
	if err != nil {
		return SettingsStruct{}, fmt.Errorf("error parsing settings.json file: %w", err)
	}

	err = file.Close()
	if err != nil {
		return SettingsStruct{}, fmt.Errorf("error closing settings.json file: %w", err)
	}

	return
}

func ParseActions() (actions []Action, err error) {
	byteValue, err := os.Open("./json/actions.json")
	if err != nil {
		return nil, fmt.Errorf("error reading actions.json file: %w", err)
	}

	actions, err = ParseActionsBytes(byteValue)
	if err != nil {
		return nil, fmt.Errorf("error parsing actions.json file: %w", err)
	}

	return
}

func ParseActionsBytes(reader io.Reader) (actions []Action, err error) {
	var mapData map[string]interface{}

	err = json.NewDecoder(reader).Decode(&mapData)
	if err != nil {
		return
	}

	for key, value := range mapData {
		action := GetActionFromType(key, value.(map[string]interface{})["type"].(string))
		err = mapstructure.Decode(value, &action)
		if err != nil {
			return
		}

		actions = append(actions, action)
	}

	slices.SortFunc(actions, func(a, b Action) int { return strings.Compare(a.Name, b.Name) })
	return
}

func ParseMemeList() (memes []Meme, err error) {
	byteValue, err := os.Open("./json/memes.json")
	if err != nil {
		return nil, fmt.Errorf("error reading memes.json file: %w", err)
	}

	var mapData map[string]string
	err = json.NewDecoder(byteValue).Decode(&mapData)
	if err != nil {
		return nil, fmt.Errorf("error parsing memes.json file: %w", err)
	}

	for key, value := range mapData {
		meme := Meme{Name: key, Text: value}
		memes = append(memes, meme)
	}

	return
}

func ParseOrCreateGroups() (GroupsStruct, error) {
	groups := make(GroupsStruct)

	byteValue, err := os.ReadFile(groupsPath)
	if errors.Is(err, os.ErrNotExist) {
		return groups, nil
	} else if err != nil {
		return nil, fmt.Errorf("error reading groups.json file: %w", err)
	}

	err = json.Unmarshal(byteValue, &groups)
	if err != nil {
		return nil, fmt.Errorf("error parsing groups.json file: %w", err)
	}

	if groups == nil {
		groups = make(GroupsStruct)
	}

	return groups, nil
}

func SaveGroups(groups GroupsStruct) error { return utils.WriteJSONFile(groupsPath, groups) }
