func findDegrees(matrix [][]int) []int {
        n := len(matrix)
    ans := make([]int, n)

    for i := 0; i < n; i++ {
        degree := 0
        for j := 0; j < n; j++ {
            degree += matrix[i][j]
        }
        ans[i] = degree
    }

    return ans
}