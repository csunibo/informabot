package model

import (
	"fmt"
	"testing"
	"time"

	tgbotapi "github.com/musianisamuele/telegram-bot-api"
)

func TestGetTimetableCoursesRows(t *testing.T) {
	var timetables = map[string]Timetable{
		"l_informatica": {
			Course: "Informatica",
			Type:   "laurea",
			Name:   "informatica",
		},
		"lm_informatica_software_techniques": {
			Course:     "Informatica Magistrale - Tecniche del software",
			Type:       "magistrale",
			Name:       "informatica",
			Curriculum: "A58-000",
		},
	}

	type args struct {
		data map[string]Timetable
	}
	tests := []struct {
		name string
		args args
		want InlineKeyboardRows
	}{
		{
			name: "All the timetables from the map",
			args: args{data: timetables},
			want: InlineKeyboardRows{
				tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Informatica", "lectures_l_informatica")),
				tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Informatica Magistrale - Tecniche del software", "lectures_lm_informatica_software_techniques")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got InlineKeyboardRows = GetTimetableCoursesRows(&tt.args.data)
			if len(got) != len(tt.want) {
				t.Errorf("GetTimetableCoursesRows() = %v, want %v", got, tt.want)
			} else {
				for i, v := range got {
					for j, w := range v {
						if w.Text != tt.want[i][j].Text || *w.CallbackData != *tt.want[i][j].CallbackData {
							t.Errorf("GetTimetableCoursesRows() = %v, want %v", w, tt.want[i][j])
						}
					}
				}
			}
		})
	}

}

func TestChooseTimetableDay(t *testing.T) {
	dt := time.Now()
	var weekdays = [7]string{
		"Domenica", "Lunedì", "Martedì", "Mercoledì", "Giovedì", "Venerdì", "Sabato",
	}
	var months = [12]string{
		"Dicembre", "Gennaio", "Febbraio", "Marzo", "Aprile", "Maggio", "Giugno", "Luglio", "Agosto", "Settembre", "Ottobre", "Novembre",
	}

	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want InlineKeyboardRows
	}{
		{
			name: "Get lectures for the week",
			args: args{data: "lectures_lm_informatica_software_techniques"},
			want: make([][]tgbotapi.InlineKeyboardButton, 7),
		},
	}

	for day := 0; day < 7; day++ {
		tests[0].want[day] = tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%s %d %s", weekdays[dt.Weekday()], dt.Day(), months[dt.Month()]), fmt.Sprintf("%s_day_%d", "lectures_lm_informatica_software_techniques", dt.Unix())))
		dt = dt.AddDate(0, 0, 1)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got InlineKeyboardRows = ChooseTimetableDay(tt.args.data)
			if len(got) != len(tt.want) {
				t.Errorf("ChooseTimetableDay() = %v, want %v", got, tt.want)
			} else {
				for i, v := range got {
					for j, w := range v {
						if w.Text != tt.want[i][j].Text || *w.CallbackData != *tt.want[i][j].CallbackData {
							t.Errorf("ChooseTimetableDay() = %v, want %v", w, tt.want[i][j])
						}
					}
				}
			}
		})
	}
}

func TestGetLectureYears(t *testing.T) {
	type args struct {
		data [2]string
	}
	tests := []struct {
		name string
		args args
		want InlineKeyboardRows
	}{
		{
			name: "Get rows for bachelor's degree",
			args: args{data: [2]string{"lectures_l_informatica", "Informatica"}},
			want: InlineKeyboardRows{
				tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Informatica: 1^ anno", "lectures_l_informatica_y_1")),
				tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Informatica: 2^ anno", "lectures_l_informatica_y_2")),
				tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Informatica: 3^ anno", "lectures_l_informatica_y_3")),
			},
		},
		{
			name: "Get rows for master's degree",
			args: args{data: [2]string{"lectures_lm_informatica_software_techniques", "Informatica Magistrale"}},
			want: InlineKeyboardRows{
				tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Informatica Magistrale: 1^ anno", "lectures_lm_informatica_software_techniques_y_1")),
				tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Informatica Magistrale: 2^ anno", "lectures_lm_informatica_software_techniques_y_2")),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got InlineKeyboardRows = GetLectureYears(tt.args.data[0], tt.args.data[1])
			if len(got) != len(tt.want) {
				t.Errorf("GetLectureYears() = %v, want %v", got, tt.want)
			} else {
				for i, v := range got {
					for j, w := range v {
						if w.Text != tt.want[i][j].Text || *w.CallbackData != *tt.want[i][j].CallbackData {
							t.Errorf("GetLectureYears() = %v, want %v", w, tt.want[i][j])
						}
					}
				}
			}
		})
	}
}
