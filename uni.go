package main

import (
	"encoding/json"
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

func Uni() {
	resp, err := http.Get("https://corsi.unibo.it/laurea/informatica/orario-lezioni/@@orario_reale_json?anno=1&start=2023-03-13&end=2023-03-13")
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

	log.Println(result)
}
