func relativeSortArray(arr1 []int, arr2 []int) []int {
 m := map[int]int{}

	for i, v := range arr2 {
		m[v] = i
	}

	sort.Slice(arr1, func(i, j int) bool {
		x, ok1 := m[arr1[i]]
		y, ok2 := m[arr1[j]]

		if ok1 && ok2 { return x < y }
		if ok1 { return true }
		if ok2 { return false }
		return arr1[i] < arr1[j]
	})

	return arr1
}