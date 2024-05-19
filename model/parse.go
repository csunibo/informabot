package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/mapstructure"
	"golang.org/x/exp/slices"

	"github.com/csunibo/informabot/utils"
)

const (
	jsonPath            = "./json/"
	groupsFile          = "groups.json"
	configSubpath       = "config/"
	degreesFile         = "degrees.json"
	teachingsFile       = "teachings.json"
	timetablesFile      = "timetables.json"
	maintainersFile     = "maintainers.json"
	representativesFile = "representatives.json"
)

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

const COMMAND_MAX_LENGTH = 32

func commandNameFromString(s string) string {
	s = utils.ToSnakeCase(s)
	if len(s) > COMMAND_MAX_LENGTH {
		return s[:COMMAND_MAX_LENGTH]
	}
	return s
}

func commandNameFromTeaching(t Teaching) string {
	return commandNameFromString(t.Url)
}

func commandNameFromDegree(d Degree) string {
	return commandNameFromString(d.Id)
}

func commandNamesFromStrings(strings []string) {
	for i, s := range strings {
		strings[i] = commandNameFromString(s)
	}
}

func ParseTeachings() (teachings map[string]Teaching, err error) {
	filepath := filepath.Join(jsonPath, configSubpath, teachingsFile)
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("error reading %s file: %w", filepath, err)
	}

	var teachingsArray []Teaching
	err = json.NewDecoder(file).Decode(&teachingsArray)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s file: %w", filepath, err)
	}
	teachings = make(map[string]Teaching, len(teachingsArray))
	for _, t := range teachingsArray {
		teachings[commandNameFromTeaching(t)] = t
	}
	return
}

func ParseDegrees() (degrees map[string]Degree, err error) {
	filepath := filepath.Join(jsonPath, configSubpath, degreesFile)
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("error reading %s file: %w", degreesFile, err)
	}
	var degreesArray []Degree
	err = json.NewDecoder(file).Decode(&degreesArray)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s file: %w", degreesFile, err)
	}
	for _, d := range degreesArray {
		for _, y := range d.Years {
			t := y.Teachings
			commandNamesFromStrings(t.Mandatory)
			commandNamesFromStrings(t.Electives)
		}
	}
	degrees = make(map[string]Degree, len(degreesArray))
	for _, d := range degreesArray {
		degrees[commandNameFromDegree(d)] = d
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

	filepath := filepath.Join(jsonPath, groupsFile)
	byteValue, err := os.ReadFile(filepath)
	if errors.Is(err, os.ErrNotExist) {
		return groups, nil
	} else if err != nil {
		return nil, fmt.Errorf("error reading %s file: %w", filepath, err)
	}

	err = json.Unmarshal(byteValue, &groups)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s file: %w", filepath, err)
	}

	if groups == nil {
		groups = make(GroupsStruct)
	}

	return groups, nil
}

func SaveGroups(groups GroupsStruct) error {
	filepath := filepath.Join(jsonPath, groupsFile)
	return utils.WriteJSONFile(filepath, groups)
}

func ParseTimetables() (timetables map[string]Timetable, err error) {
	filepath := filepath.Join(jsonPath, configSubpath, timetablesFile)
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("error reading %s file: %w", timetablesFile, err)
	}

	var mapData map[string]Timetable

	err = json.NewDecoder(file).Decode(&mapData)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s file: %w", filepath, err)
	}

	timetables = mapData
	return
}

func ParseMaintainers() (maintainer []Maintainer, err error) {
	filepath := filepath.Join(jsonPath, configSubpath, maintainersFile)
	file, err := os.ReadFile(filepath)
	if errors.Is(err, os.ErrNotExist) {
		return maintainer, fmt.Errorf("%s does not exist", maintainersFile)
	} else if err != nil {
		return nil, fmt.Errorf("error reading %s file: %w", maintainersFile, err)
	}

	var projects []struct {
		Name        string       `json:"project"`
		Maintainers []Maintainer `json:"maintainers"`
	}

	err = json.Unmarshal(file, &projects)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s file: %w", maintainersFile, err)
	}

	for _, p := range projects {
		if p.Name == "informabot" {
			return p.Maintainers, nil
		}
	}

	return nil, fmt.Errorf("couldn't found informabot projects after parsing %s", maintainersFile)
}

func ParseRepresentatives() (map[string]Representative, error) {
	representatives := make(map[string]Representative)

	filepath := filepath.Join(jsonPath, configSubpath, representativesFile)
	byteValue, err := os.ReadFile(filepath)
	if errors.Is(err, os.ErrNotExist) {
		return representatives, nil
	} else if err != nil {
		return nil, fmt.Errorf("error reading %s file: %w", filepath, err)
	}

	err = json.Unmarshal(byteValue, &representatives)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s file: %w", filepath, err)
	}

	if representatives == nil {
		representatives = make(map[string]Representative)
	}

	return representatives, nil
}
