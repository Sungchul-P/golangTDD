package math

import (
	"errors"
)

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func Div(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("Can't divide by zero")
	}
	return a / b, nil
}

func StrRept(s string, count int) string {
	b := make([]byte, len(s)*count)
	bp := copy(b, s)
	for bp < len(b) {
		copy(b[bp:], b[bp:])
		bp *= 2
	}
	return string(b)
}
