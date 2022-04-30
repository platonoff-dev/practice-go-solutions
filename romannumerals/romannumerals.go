package romannumerals

type Pair struct {
	Roman string
	Value int
}

func Encode(n int) (string, bool) {
	if n <= 0 {
		return "", false
	}

	romanMappings := []Pair{
		{Roman: "M", Value: 1000},
		{Roman: "CM", Value: 900},
		{Roman: "D", Value: 500},
		{Roman: "CD", Value: 400},
		{Roman: "C", Value: 100},
		{Roman: "XC", Value: 90},
		{Roman: "L", Value: 50},
		{Roman: "XL", Value: 40},
		{Roman: "X", Value: 10},
		{Roman: "IX", Value: 9},
		{Roman: "V", Value: 5},
		{Roman: "IV", Value: 4},
		{Roman: "I", Value: 1},
	}

	result := ""
	for n > 0 {
		for _, m := range romanMappings {
			if m.Value <= n {
				n -= m.Value
				result += m.Roman
				break
			}
		}
	}

	return result, true
}

func Decode(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	romanMappings := []Pair{
		// paired values
		{Roman: "IV", Value: 4},
		{Roman: "IX", Value: 9},
		{Roman: "XL", Value: 40},
		{Roman: "XC", Value: 90},
		{Roman: "CD", Value: 400},
		{Roman: "CM", Value: 900},
		// single values
		{Roman: "I", Value: 1},
		{Roman: "V", Value: 5},
		{Roman: "X", Value: 10},
		{Roman: "L", Value: 50},
		{Roman: "C", Value: 100},
		{Roman: "D", Value: 500},
		{Roman: "M", Value: 1000},
	}

	number := 0
	pos := 0

	for pos < len(s) {
		found := false

		for _, m := range romanMappings {
			roman := m.Roman
			value := m.Value

			// Get substring same length as roman number
			substr := ""
			if pos+len(roman)-1 < len(s) {
				substr = s[pos : pos+len(roman)]
			} else {
				continue
			}

			if substr == roman {
				found = true
				pos += len(roman)
				number += value
				break
			}
		}

		if !found {
			return 0, false
		}
	}

	return number, true
}
