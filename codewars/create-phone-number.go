package kata

func numbersToString(numbers [10]uint) string {
	ans := ""

	for _, num := range numbers {
		ans += string(num + '0')
	}
	return ans
}

func CreatePhoneNumber(numbers [10]uint) string {
	numString := numbersToString(numbers)
	return "(" + string(numString[:3]) + ") " + numString[3:6] + "-" + numString[6:]
}
