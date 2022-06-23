package helpers

import (
	"strings"
)

// * Using modified Ratcliff-Obershelp similarity function
func CompareStrings(s string, t string) int {
	if len(s) == 1 || len(t) == 1 {
		return 0
	}

	maxSub := string(getMaxLengthCommonString(s, t))
	sIdx := strings.Index(s, maxSub)
	tIdx := strings.Index(t, maxSub)

	sLeft := string(s[:sIdx])
	tLeft := string(t[:tIdx])

	sRight := string(s[sIdx+(len(maxSub)-1):])
	tRight := string(t[tIdx+(len(maxSub)-1):])

	return len(maxSub) + CompareStrings(sLeft, tLeft) + CompareStrings(sRight, tRight)
}

func getMaxLengthCommonString(str1, str2 string) string {
	answer := ""

	len1, len2 := len(str1), len(str2)

	for i := 0; i < len1; i++ {
		match := ""
		for j := 0; i < len2; i++ {
			if i+j < len1 && str1[i+j] == str2[j] {
				match += string(str2[j])
			} else {
				if len(match) > len(answer) {
					answer = match
				}
			}
		}
	}
	return answer
}
