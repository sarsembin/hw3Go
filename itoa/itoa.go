package itoa

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var number []byte
	answer := ""
	if n < 0 {
		answer += "-"
		n *= -1
	}
	for n > 0 {
		digit := n % 10
		number = append(number, byte(digit + '0'))
		n /= 10
	}
	for i := 0; i < len(number) / 2; i++ {
		number[i], number[len(number) - 1 - i] = number[len(number) - 1 - i], number[i]
	}
	return answer + string(number)
}
