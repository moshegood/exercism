package luhn

func Valid(input string) bool {
	numbers := []int{}
	for _, char := range input {
		if char == ' ' {
			continue
		}
		if char < '0' || char > '9' {
			return false
		}
		numbers = append(numbers, int(char-'0'))
	}

	length := len(numbers)
	if length < 2 {
		return false
	}

	sum := 0
	for i, value := range numbers {
		if (i+length)%2 == 0 {
			value *= 2
			if value > 9 {
				value -= 9
			}
		}
		sum += value
	}
	return sum%10 == 0
}
