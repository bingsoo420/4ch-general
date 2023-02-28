package main_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	. "github.com/bingsoo420/4ch-general"
)

func TestItParseValidSubjects(t *testing.T) {
	tests := []struct {
		subject string
		valid   bool
		want    string
	}{
		{"/wdg/", true, "wdg"},
		{"xdg", false, ""},
		{"Anonymous", false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.subject, func(t *testing.T) {
			isValid, extracted := ParseSubject(tt.subject)

			if isValid != tt.valid || extracted != tt.want {
				t.Errorf("got [%t, %s], want [%t, %s]", isValid, extracted, tt.valid, tt.want)
			}
		})
	}
}

func TestItParseGeneralsFromFixtures(t *testing.T) {
	tests := []struct {
		board  string
		length int
	}{
		{"g", 24},
		{"fit", 7},
		{"biz", 14},
	}

	for _, tt := range tests {
		t.Run(tt.board, func(t *testing.T) {
			dat, err := os.ReadFile(fmt.Sprintf("./fixtures/%s_catalog.json", tt.board))
			if err != nil {
				panic(err)
			}

			var catalog Catalog

			if err := json.Unmarshal(dat, &catalog); err != nil {
				panic(err)
			}

			generals := BuildGenerals("g", catalog)

			if len(generals) != tt.length {
				t.Errorf("got %d, want %d", len(generals), tt.length)
			}
		})
	}
}
