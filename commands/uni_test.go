package commands

import (
	"strings"
	"testing"
	"time"
)

func TestLezioniTime_Format(t *testing.T) {
	a := time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC)
	ti := LezioniTime(a)

	type args struct {
		format string
	}
	tests := []struct {
		name string
		tr   *LezioniTime
		args args
		want string
	}{
		{
			name: "Date formatting",
			tr:   &ti,
			args: args{
				format: "2006-01-02T15:04:05",
			},
			want: "2023-01-01T12:00:00",
		},
		{
			name: "Empty date",
			tr:   &LezioniTime{},
			args: args{
				format: "2006-01-02T15:04:05",
			},
			want: "0001-01-01T00:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := tt.tr
			if got := tr.Format(tt.args.format); got != tt.want {
				t.Errorf("LezioniTime.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLezioniTime_UnmarshalJSON(t *testing.T) {
	var time LezioniTime
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		tr      *LezioniTime
		args    args
		wantErr bool
	}{
		{
			name: "Not an error",
			tr:   &time,
			args: args{
				data: []byte{34, 50, 48, 50, 51, 45, 48, 51, 45, 49, 51, 84, 49, 50, 58, 48, 48, 58, 48, 48, 34},
			},
			wantErr: false,
		},
		{
			name: "Error #1",
			tr:   &time,
			args: args{
				data: []byte("A"),
			},
			wantErr: true,
		},
		{
			name: "Error #2",
			tr:   &time,
			args: args{
				data: []byte{},
			},
			wantErr: true,
		},
		{
			name: "Error #2",
			tr:   &time,
			args: args{
				data: []byte{34},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := tt.tr
			if err := tr.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("LezioniTime.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetTimeTable(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Weekend",
			args: args{
				url: "https://corsi.unibo.it/laurea/informatica/orario-lezioni/@@orario_reale_json?anno=1&start=2023-03-11&end=2023-03-11",
			},
			want: "",
		},
		{
			name: "Weekday",
			args: args{
				url: "https://corsi.unibo.it/laurea/informatica/orario-lezioni/@@orario_reale_json?anno=1&start=2023-03-13&end=2023-03-13",
			},
			want: `  🕘 <b><a href="">ALGORITMI E STRUTTURE DI DATI / (2) Modulo 2</a></b>
					12:00 - 14:00
					🏢 AULA MAGNA - Piano Terra
					📍 Via Filippo Re, 10 - Bologna
					🕘 <b><a href="">ANALISI MATEMATICA / (2) Modulo 2</a></b>
					15:00 - 18:00
					🏢 AULA MAGNA - Piano Terra
					📍 Via Filippo Re, 10 - Bologna
					`,
		},
		{
			name: "Not a valid url",
			args: args{
				url: "https://example.com",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTimeTable(tt.args.url)
			got = strings.ReplaceAll(got, " ", "")
			want := strings.ReplaceAll(tt.want, " ", "")
			want = strings.ReplaceAll(want, "\t", "")
			if got != want {
				t.Errorf("GetTimeTable() = %v, want %vs", got, want)
			}
		})
	}
}
