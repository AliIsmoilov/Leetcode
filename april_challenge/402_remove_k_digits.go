package aprilchallenge

func RemoveKdigits(num string, k int) string {
	res := make([]rune, 0)

	for _, ch := range num {
		for len(res) > 0 && res[len(res)-1] > ch && k > 0 {
			res = res[:len(res)-1]
			k--
		}

		if len(res) > 0 || string(ch) != "0" {
			res = append(res, ch)
		}
	}

	for len(res) > 0 && k > 0 {
		res = res[:len(res)-1]
		k--
	}

	if len(res) == 0 {
		return "0"
	}

	return string(res)
}
