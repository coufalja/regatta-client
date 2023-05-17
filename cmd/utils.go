package cmd

func findNextString(str string) string {
	// Convert string to byte slice for mutation
	bytes := []byte(str)

	// Start from the last character and increment its byte value
	i := len(bytes) - 1
	for i >= 0 {
		if bytes[i] < 255 {
			bytes[i]++
			break
		}
		bytes[i] = 0
		i--
	}

	return string(bytes)
}
