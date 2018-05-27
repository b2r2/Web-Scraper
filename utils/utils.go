package utils

import (
	"strings"
)

func IsURL(url string) bool {
	line := strings.Split(url, "://")
	if line[0] == "http" || line[0] == "https" {
		return true
	} else {
		return false
	}
}
