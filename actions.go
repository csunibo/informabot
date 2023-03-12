package main

type DataInterface interface{}

func GetActionFromType(commandType string) Action {
	var data DataInterface
	switch commandType {
	case "message":
		data = MessageData{}
	case "help":
		data = HelpData{}
	case "aggiorna":
		data = AggiornaData{}
	case "lookingFor":
		data = LookingForData{}
	case "notLookingFor":
		data = NotLookingForData{}
	case "yearly":
		data = YearlyData{}
	case "todayLecture":
		data = TodayLectureData{}
	case "tomorrowLecture":
		data = TomorrowLectureData{}
	case "scelta":
		data = SceltaData{}
	case "course":
		data = CourseData{}
	default:
		data = MessageData{}
	}

	return Action{
		Name: "UNINITIALIZED",
		Type: commandType,
		Data: data,
	}
}

type Action struct {
	Name string
	Type string        `json:"type"`
	Data DataInterface `json:"data"`
}

type MessageData struct {
	Text string `json:"text"`
}

type HelpData struct {
	Description string `json:"description"`
}

type AggiornaData struct {
	Description string `json:"description"`
	NoYear      string `json:"noYear"`
	NoMod       string `json:"noMod"`
	Started     string `json:"started"`
	Ended       string `json:"ended"`
	Failed      string `json:"failed"`
}

type LookingForData struct {
	Description  string `json:"description"`
	SingularText string `json:"singularText"`
	PluralText   string `json:"pluralText"`
	ChatError    string `json:"chatError"`
}

type NotLookingForData struct {
	Description   string `json:"description"`
	Text          string `json:"text"`
	ChatError     string `json:"chatError"`
	NotFoundError string `json:"notFoundError"`
}

type YearlyData struct {
	Description string `json:"description"`
	Command     string `json:"command"`
	NoYear      string `json:"noYear"`
}

type TodayLectureData struct {
	Description  string `json:"description"`
	Url          string `json:"url"`
	Title        string `json:"title"`
	FallbackText string `json:"fallbackText"`
}

type TomorrowLectureData TodayLectureData

type SceltaData struct {
	Description string     `json:"description"`
	Header      string     `json:"header"`
	Template    string     `json:"template"`
	Items       [][]string `json:"items"`
}

type CourseData struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Virtuale     string   `json:"virtuale"`
	Teams        string   `json:"teams"`
	Website      string   `json:"website"`
	Professors   []string `json:"professors"`
	TelegramLink string   `json:"telegram"`
}
