func threeSum(nums []int) [][]int {
    	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { continue }

		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]

			if s == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] { l++ }
				for l < r && nums[r] == nums[r-1] { r-- }
				l++; r--
			} else if s < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}