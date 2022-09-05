package page_model_test_tools

import (
	"strings"
)

func calcNewShift(oldRows []string, shift int, newRows []string) int {
	if len(oldRows) == len(newRows) {
		return shift
	}

	anchorWordNumber := 0
	if shift != 0 {
		for i := 0; i <= shift; i++ {
			anchorWordNumber += countWords(oldRows[i])
		}
		anchorWordNumber++ // add next after shift word

		counter := 0
		for shift = 0; shift < len(newRows); shift++ {
			counter += countWords(newRows[shift])
			if anchorWordNumber <= counter {
				shift--
				break
			}
		}
	}

	return shift
}

func countWords(line string) (count int) {
	words := strings.Split(line, " ")
	for _, w := range words {
		if len(w) == 0 || w == " " {
			continue
		}
		count++
	}
	if count == 0 {
		return 1
	}
	return count
}
