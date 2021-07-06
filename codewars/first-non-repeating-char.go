package kata

func FirstNonRepeating(str string) string {
	check := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		if 'A' <= str[i] && str[i] <= 'Z' {
			check[str[i]-'A'+'a']++
		} else {
			check[str[i]]++
		}
	}

	for i := 0; i < len(str); i++ {
		char := str[i]
		if 'A' <= str[i] && str[i] <= 'Z' {
			char = char - 'A' + 'a'
		}
		if check[char] == 1 {
			return string(str[i])
		}
	}
	return ""
}
