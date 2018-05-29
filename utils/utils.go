package utils

import (
	"strings"
)

func IsCorrectURL(url string) bool {
	correctUrl := []string{
		"medium",
		"telegra",
	}
	state := false
	line := strings.Split(url, "://")
	if line[0] == "http" || line[0] == "https" {
		line = strings.Split(line[1], ".")
		if line[0] == correctUrl[0] || line[0] == correctUrl[1] {
			state = true
		}
	}
	return state
}

func GetDomain(url string) string {
	line := strings.Split(url, "://")
	line = strings.Split(line[1], ".")
	return line[0]
}
