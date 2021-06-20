package kata

func Solution(str string) []string {
	if len(str)%2 == 1 {
		str += "_"
	}
	ans := make([]string, 0)
	for i := 0; i < len(str); i += 2 {
		ans = append(ans, str[i:i+2])
	}
	return ans
}
