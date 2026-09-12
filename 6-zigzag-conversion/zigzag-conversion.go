func convert(s string, numRows int) string {
    if numRows == 1 {
		return s
	}

	n := len(s)
	cycleLen := 2 * (numRows - 1)
	result := make([]byte, 0, n)

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycleLen {
			result = append(result, s[j+i])

			if i != 0 && i != numRows-1 && j+cycleLen-i < n {
				result = append(result, s[j+cycleLen-i])
			}
		}
	}

	return string(result)
}