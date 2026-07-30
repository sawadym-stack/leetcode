func equalFrequency(word string) bool {
    freq := make([]int, 26)

	for _, ch := range word {
		freq[ch-'a']++
	}

	for i := 0; i < 26; i++ {
		if freq[i] == 0 {
			continue
		}

		freq[i]--

		target := 0
		ok := true

		for _, f := range freq {
			if f == 0 {
				continue
			}
			if target == 0 {
				target = f
			} else if f != target {
				ok = false
				break
			}
		}

		freq[i]++

		if ok {
			return true
		}
	}

	return false
}