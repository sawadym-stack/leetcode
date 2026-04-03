func isAnagram(s string, t string) bool {
    // if len(s) != len(t) { return false }

	// m := map[byte]int{}
	// for i := range s {
	// 	m[s[i]]++
	// 	m[t[i]]--
	// }
	// for _, v := range m {
	// 	if v != 0 { return false }
	// }
	// return true
    if len(s) != len(t) { return false }

	a, b := []byte(s), []byte(t)
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })

	return string(a) == string(b)
}