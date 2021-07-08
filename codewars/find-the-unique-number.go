package kata

func chooseThree(arr []float32) []float32 {
	ret := make([]float32, 3)
	ret[0] = arr[0]
	lastIdx := len(arr) - 1
	for i := 0; i < len(arr); i++ {
		if arr[i] != arr[0] {
			ret[1] = arr[i]
			if i == lastIdx {
				ret[2] = arr[i-1]
			} else {
				ret[2] = arr[lastIdx]
			}
		}
	}
	return ret
}

func FindUniq(arr []float32) float32 {

	nums := chooseThree(arr)
	if nums[0] == nums[1] {
		return nums[2]
	} else if nums[1] == nums[2] {
		return nums[0]
	} else {
		return nums[1]
	}
}
