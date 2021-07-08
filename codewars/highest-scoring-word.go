package kata

import (
	"strings"
)

func scoreWord(word string) (score int) {
	score = 0
	for i := 0; i < len(word); i++ {
		score += int(word[i]-'a') + 1
	}
	return
}

func High(s string) string {
	slice := strings.Split(s, " ")
	mxScore := 0
	ans := ""
	for _, str := range slice {
		now := scoreWord(str)
		if mxScore < now {
			mxScore = now
			ans = str
		}
	}
	return ans
}
