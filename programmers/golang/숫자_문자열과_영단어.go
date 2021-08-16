package main

import (
	"strconv"
	"strings"
)

var number = [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func solution(s string) int {
	ans := s
	for i := 0; i < 10; i++ {
		ans = strings.Replace(ans, number[i], string('0'+i), -1)
	}
	result, _ := strconv.Atoi(ans)
	return result
}
