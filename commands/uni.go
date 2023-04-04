package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"
)

const TIMEFORMAT = "2006-01-02T15:04:05"

type Aula struct {
	Edificio  string `json:"des_edificio"`
	Piano     string `json:"des_piano"`
	Indirizzo string `json:"des_indirizzo"`
}

type LezioniTime time.Time

type OrarioLezioni struct {
	Title     string      `json:"title"`
	Time      string      `json:"time"`
	Aule      []Aula      `json:"aule"`
	Teams     string      `json:"teams"`
	StartTime LezioniTime `json:"start"`
	EndTime   LezioniTime `json:"end"`
}

func (t *LezioniTime) Format(format string) string {
	return (*time.Time)(t).Format(format)
}

func (t *LezioniTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsedTime, err := time.Parse(TIMEFORMAT, s)
	if err != nil {
		return err
	}
	*t = LezioniTime(parsedTime)
	return nil
}

// GetTimeTable returns the timetable for the Unibo url
// returns empty string if there are no lessons.
func GetTimeTable(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error getting json when requesting orario lezioni: %s\n", err)
	}
	defer resp.Body.Close()

	result := []OrarioLezioni{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &result)

	sort.Slice(result, func(i, j int) bool {
		return (*time.Time)(&result[i].StartTime).Before((time.Time)(result[j].StartTime))
	})

	var message string = ""
	for _, lezione := range result {
		message += fmt.Sprintf(`  🕘 <b><a href="%s">%s</a></b>`, lezione.Teams, lezione.Title) +
			"\n" + lezione.Time + "\n"
		if len(lezione.Aule) > 0 {
			message += fmt.Sprintf("  🏢 %s - %s\n", lezione.Aule[0].Edificio, lezione.Aule[0].Piano)
			message += fmt.Sprintf("  📍 %s\n", lezione.Aule[0].Indirizzo)
		}
	}

	return message
}
