package piscine

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	var cleaned []rune
	for _, ch := range str {
		if ch != ' ' {
			cleaned = append(cleaned, ch)
		}
	}

	if len(cleaned) < 5 {
		return "Invalid Input\n"
	}

	var result []rune
	count := 0

	for _, ch := range cleaned {
		count++

		if count == 6 {
			count = 0
			continue
		}

		result = append(result, ch)

		if count == 5 {
			result = append(result, ' ')
			count = 0
		}
	}

	result = append(result, '\n')
	return string(result)
}
