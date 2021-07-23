package kata

import "strconv"

func SumOfIntegersInString(strng string) int {
	score := 0
	num := make([]byte, 0)
	for i := 0; i < len(strng); i++ {
		if strng[i] >= '0' && strng[i] <= '9' {
			num = append(num, strng[i])
		} else {
			if len(num) > 0 {
				s, _ := strconv.Atoi(string(num))
				score += s
				num = num[:0]
			}
		}
	}
	if len(num) > 0 {
		s, _ := strconv.Atoi(string(num))
		score += s
		num = num[:0]
	}
	return score
}
