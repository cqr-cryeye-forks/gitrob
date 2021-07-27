package core

import (
	"testing"
)

func TestParseLogin(t *testing.T) {
	// https://github.com/user/repo.git
	// git@github.com:username/repo.git

	testTable := []struct {
		input    string
		expected string
	}{
		{
			input:    "username",
			expected: "username",
		},
		{
			input:    "user-name",
			expected: "user-name",
		},
		{
			input:    "username-",
			expected: "",
		},
		{
			input:    "user-name/",
			expected: "user-name",
		},
		{
			input:    "username#",
			expected: "",
		},
		{
			input:    "-user-name-",
			expected: "",
		},
		{
			input:    "u",
			expected: "u",
		},
		{
			input:    "_",
			expected: "",
		},
		{
			input:    "git@github.com:username/repo.git",
			expected: "username",
		},
		{
			input:    "https://github.com/username/repo.git",
			expected: "username",
		},
		{
			input:    "https://github.com/user-name/repo.git",
			expected: "user-name",
		},
		{
			input:    "https://github.com/usern@ame/repo.git",
			expected: "",
		},
		{
			input:    "https://github.com/username",
			expected: "username",
		},
		{
			input:    "https://github.com/username/",
			expected: "username",
		},
	}

	for _, testCase := range testTable {
		result := ParseLogin(testCase.input)
		if result != testCase.expected {
			t.Errorf("Unexpected result! Ecpected <%v> - got <%v>", testCase.expected, result)
		}
	}
}
