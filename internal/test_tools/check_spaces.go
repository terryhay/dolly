package test_tools

import (
	"fmt"
	"regexp"
	"strings"
)

// CheckSpaces checks if text contains too much new lines or spaces in a row
func CheckSpaces(text string) (res bool, msg string) {

	if regexp.MustCompilePOSIX(" {2}").MatchString(text) {
		return false, fmt.Sprintf("text \"%s\" contains too much spaces in a row", text)
	}

	emptyStringCounter := 0
	textRows := strings.Split(text, "\n")

	for i := 0; i < len(textRows)-1; i++ {
		if strings.TrimSpace(textRows[i]) == "" {
			emptyStringCounter++
			if emptyStringCounter == 2 {
				return false, fmt.Sprintf("text \"%s\" contains too much new lines in a row", text)
			}
			continue
		}

		emptyStringCounter = 0
	}

	return true, ""
}
