package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseSubject(subject string) (bool, string) {
	re := regexp.MustCompile(`\/(\w+)\/`)

	find := re.Find([]byte(subject))

	return len(find) > 0, strings.Trim(string(find), "/")
}

func BuildGenerals(board string, catalog []Page) map[string]string {
	generals := make(map[string]string)

	for _, page := range catalog {
		for _, thread := range page.Threads {
			if isSubject, subject := ParseSubject(thread.Sub); isSubject {
				subject := strings.ToLower(subject)
				generals[subject] = fmt.Sprintf("https://boards.4chan.org/%s/thread/%d", board, thread.No)
			}
		}
	}

	return generals
}
