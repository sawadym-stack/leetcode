func missingInteger(nums []int) int {
      sum := nums[0]
    
    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1]+1 {
            sum += nums[i]
        } else {
            break
        }
    }

    seen := make(map[int]bool)
    for _, num := range nums {
        seen[num] = true
    }

    x := sum
    for {
        if !seen[x] {
            return x
        }
        x++
    }
}