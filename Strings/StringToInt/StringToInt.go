package StringToInt

import (
	"math"
	"strings"
	"unicode"
)

func myAtoi(s string) int {
	val := int64(0)
	isNegative := false
	s = strings.TrimSpace(s)

	exponent := uint64(1)
	counter := 0

	for i, c := range s {
		if !unicode.IsDigit(rune(c)) {
			if c == '-' && i == 0 {
				isNegative = true
				s = strings.TrimLeft(s, "-")
				continue
			} else if c == '+' && i == 0 {
				s = strings.TrimLeft(s, "+")
				continue
			} else {
				break
			}
		} else if c == '0' && exponent < 10 {
			s = strings.TrimLeft(s, "0")
			continue
		}
		counter++
		if counter > 10 {
			if isNegative {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}

		exponent *= 10
	}

	if isNegative {
		if exponent/10 > math.MaxInt32+1 {
			return math.MinInt32
		}
	}

	if exponent/10 > math.MaxInt32 {
		return math.MaxInt32
	}

	for _, c := range s {
		exponent /= 10
		val += int64(c-48) * int64(exponent)
	}

	if isNegative {
		val *= -1
		if val < math.MinInt32 {
			return math.MinInt32
		}
	}

	if val > math.MaxInt32 {
		return math.MaxInt32
	}

	return int(val)
}
