func lengthOfLongestSubstring(s string) int {
	var tmpChars []string

	// explode
	inputs := strings.Split(s, "")
	stringCount := 0

	// foreach
	var parts []string
	x := 0
	for x < len(inputs) {

		// set data to be iterated
		if x == 0 {
			parts = inputs
		} else {
			parts = inputs[x:]
		}

		if len(parts) < stringCount {
			break
		}

		// iterate parts
		for _, char := range parts {

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

	fmt.Println(tmpChars)

	// return
	return stringCount
}