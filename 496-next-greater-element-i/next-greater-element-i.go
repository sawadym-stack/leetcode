func nextGreaterElement(nums1 []int, nums2 []int) []int {
    nextGreater := make(map[int]int)
	stack := []int{}

	for _, num := range nums2 {
		for len(stack) > 0 && num > stack[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextGreater[top] = num
		}
		stack = append(stack, num)
	}

	for _, num := range stack {
		nextGreater[num] = -1
	}

	result := make([]int, len(nums1))
	for i, num := range nums1 {
		result[i] = nextGreater[num]
	}

	return result
}