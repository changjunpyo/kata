package kata

func DigitalRoot(n int) int {
	if n == 0 {
		return 0
	}
	num := n
	for num >= 10 {
		tmp, next := num, 0
		for tmp/10 != 0 {
			next += tmp % 10
			tmp /= 10
		}
		next += tmp % 10
		num = next
	}
	return num
}
