package core

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var NewlineRegex = regexp.MustCompile(`\r?\n`)

// GitHubUsernameRegex UsernameRegex GitHub Username maximum is 39 characters
// May only contain alphanumeric characters or single hyphens, and cannot begin or end with a hyphen
var GitHubUsernameRegex = regexp.MustCompile(`[/:][a-zA-Z0-9-]{1,39}/`)
var UsernameRegex = regexp.MustCompile(`^([a-zA-Z0-9][a-zA-Z0-9-]{1,37})?[a-zA-Z0-9]$`)

func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func Pluralize(count int, singular string, plural string) string {
	if count == 1 {
		return singular
	}
	return plural
}

func TruncateString(str string, maxLength int) string {
	str = NewlineRegex.ReplaceAllString(str, " ")
	str = strings.TrimSpace(str)
	if len(str) > maxLength {
		str = fmt.Sprintf("%s...", str[0:maxLength])
	}
	return str
}

func ParseLogin(login string) string {
	str := GitHubUsernameRegex.FindString(login)
	if str == "" {
		str = login
	}
	str = strings.TrimPrefix(str, ":")
	str = strings.Trim(str, "/")
	if !UsernameRegex.MatchString(str) {
		return ""
	}
	return str
}
