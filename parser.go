package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseSubject(subject string) (bool, string) {
	re := regexp.MustCompile(`\/\w+\/`)

	found := re.Find([]byte(subject))

	return len(found) > 0, strings.Trim(string(found), "/")
}

func BuildGenerals(board string, catalog Catalog) map[string]string {
	generals := make(map[string]string)

	for _, page := range catalog {
		for _, thread := range page.Threads {
			if isGeneral, name := ParseSubject(thread.Sub); isGeneral {
				name := strings.ToLower(name)
				generals[name] = fmt.Sprintf("https://boards.4chan.org/%s/thread/%d", board, thread.No)
			}
		}
	}

	return generals
}
