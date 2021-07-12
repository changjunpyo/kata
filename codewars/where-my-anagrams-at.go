package kata

import "sort"

func Anagrams(word string, words []string) []string {
	str1 := []rune(word)
	sort.Slice(str1, func(i, j int) bool { return str1[i] < str1[j] })
	ans := make([]string, 0)
	for _, w := range words {
		str2 := []rune(w)
		sort.Slice(str2, func(i, j int) bool { return str2[i] < str2[j] })
		if string(str1) == string(str2) {
			ans = append(ans, w)
		}
	}
	if len(ans) == 0 {
		return nil
	}
	return ans
}
