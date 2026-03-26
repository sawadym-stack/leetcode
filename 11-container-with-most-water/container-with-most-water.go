func maxArea(height []int) int {
     l, r := 0, len(height)-1
    maxWater := 0

    for l < r {
        h := min(height[l], height[r])
        width := r - l
        area := h * width

        if area > maxWater {
            maxWater = area
        }

        if height[l] < height[r] {
            l++
        } else {
            r--
        }
    }

    return maxWater
}