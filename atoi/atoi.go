package atoi

import (
	"errors"
)

func Atoi(s string) (int, error) {
	var nanError error = errors.New("input string is not a number")
	result := 0
	sign := 1
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			if i != 0 {
				return 0, nanError
			}
			sign = -1
		} else if s[i] < '0' || s[i] > '9' {
			return 0, nanError
		} else {
			result = result * 10 + int(s[i] - '0')
		}
	}
	return result * sign, nil
}
