package kata

func duplicate_count(s1 string) int {
	//your code goes here
	ans := 0
	checkDuplicate := make([]int, 36)

	for _, v := range s1 {

		if 'A' <= v && v <= 'Z' {
			v -= ('A' - 10)
		} else if 'a' <= v && v <= 'z' {
			v -= ('a' - 10)
		} else if '0' <= v && v <= '9' {
			v -= '0'
		}
		checkDuplicate[v]++
		if checkDuplicate[v] == 2 {
			ans++
		}
	}
	return ans //Instead of returning '1', you have to return integer 'i' as answer of solution.
}
