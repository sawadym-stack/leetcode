func firstUniqChar(s string) int {
  freq := [26]int{}

    for i := 0; i < len(s); i++ {
        freq[s[i]-'a']++
    }

    for i := 0; i < len(s); i++ {
        if freq[s[i]-'a'] == 1 {
            return i
        }
    }

    return -1  
}