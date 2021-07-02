package kata

import "strconv"

func parseByPunctuation(str string) []string {
	ret := make([]string, 0)
	temp := ""
	for i := 0; i < len(str); i++ {
		if str[i] == '.' {
			ret = append(ret, string(temp))
			temp = ""
		} else {
			temp += string(str[i])
		}
	}
	ret = append(ret, temp)
	return ret
}

func Is_valid_ip(ip string) bool {

	arr := parseByPunctuation(ip)
	if len(arr) != 4 {
		return false
	}

	for i := 0; i < len(arr); i++ {
		num, err := strconv.ParseUint(arr[i], 10, 32)
		if err != nil || num > 255 {
			return false
		} else if strconv.FormatUint(num, 10) != arr[i] {
			return false
		}
	}

	return true
}
