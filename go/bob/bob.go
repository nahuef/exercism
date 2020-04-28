// Package bob returns very limited responses as he is a lackadaisical teenager.
package bob

import (
	"strings"
	"unicode"
)

func cleanInput(input string) string {
	input = strings.TrimSpace(input)

	return input
}

func isYell(remark string) bool {
	var letters = 0
	var uppercaseLetters = 0

	for _, r := range remark {
		if unicode.IsLetter(r) {
			letters++

			if unicode.IsUpper(r) {
				uppercaseLetters++
			}
		}
	}

	if letters > 0 && letters == uppercaseLetters {
		return true
	}

	return false
}

func isQuestion(remark string) bool {
	if !isEmpty(remark) &&
		strings.HasSuffix(remark, "?") {
		return true
	}

	return false
}

func isEmpty(remark string) bool {
	if len(remark) == 0 {
		return true
	}

	return false
}

// Hey returns Bob's response to a remark.
func Hey(remark string) string {
	remark = cleanInput(remark)

	var yell = isYell(remark)
	var question = isQuestion(remark)
	var empty = isEmpty(remark)

	if yell && question {
		return "Calm down, I know what I'm doing!"
	}

	if yell {
		return "Whoa, chill out!"
	}

	if question {
		return "Sure."
	}

	if empty {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
