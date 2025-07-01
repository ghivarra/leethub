func lengthOfLongestSubstring(s string) int {
	var tmpChars []string

	// explode
	inputs := strings.Split(s, "")
	stringCount := 0

	// foreach
	i := 0
	for i < len(inputs) {

		if len(inputs[i:]) < stringCount {
			break
		}

		for _, char := range inputs[i:] {
			if slices.Contains(tmpChars, char) {
				if len(tmpChars) > stringCount {
					stringCount = len(tmpChars)
				}
				tmpChars = []string{char}
				continue
			}
			tmpChars = append(tmpChars, char)
		}

		i++
	}

	fmt.Println(tmpChars)

	if len(tmpChars) > stringCount {
		stringCount = len(tmpChars)
	}

	// return
	return stringCount
}