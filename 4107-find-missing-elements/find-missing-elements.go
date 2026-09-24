func findMissingElements(nums []int) []int {
    set := make(map[int]bool)
	
	min, max := nums[0], nums[0]

	for _, num := range nums {
		set[num] = true
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	var result []int

	for i := min; i <= max; i++ {
		if !set[i] {
			result = append(result, i)
		}
	}

	return result
}