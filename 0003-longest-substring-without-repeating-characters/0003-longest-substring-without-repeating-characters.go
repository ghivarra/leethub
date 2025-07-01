func lengthOfLongestSubstring(s string) int {
	var tmpChars []string

	// explode
	inputs := strings.Split(s, "")
	stringCount := 0

	// foreach
	x := 0
	for x < len(inputs) {

		if len(inputs[x:]) < stringCount {
			break
		}

		// iterate parts
		for _, char := range inputs[x:] {

			// if there is copy
			// then empty tmpchars and break
			if slices.Contains(tmpChars, char) {
				tmpChars = []string{}
				break
			}

			// add to tmp chars
			tmpChars = append(tmpChars, char)

			// if larger than count now then add
			if len(tmpChars) > stringCount {
				stringCount = len(tmpChars)
			}
		}

		// increment x
		x++
	}

	// return
	return stringCount
}