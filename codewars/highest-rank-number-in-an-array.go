package kata

func HighestRank(nums []int) int {
	count := make(map[int]int, 0)

	for _, num := range nums {
		count[num]++
	}

	maxNum := -1
	maxCount := 0
	for key, val := range count {
		if val == maxCount && key > maxNum {
			maxNum, maxCount = key, val
		}
		if val > maxCount {
			maxNum, maxCount = key, val
		}
	}
	return maxNum
}
