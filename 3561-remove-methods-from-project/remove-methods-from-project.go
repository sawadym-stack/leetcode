func remainingMethods(n int, k int, invocations [][]int) []int {
	graph := make([][]int, n)

	for _, edge := range invocations {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
	}

	suspicious := make([]bool, n)

	var dfs func(int)
	dfs = func(node int) {
		if suspicious[node] {
			return
		}

		suspicious[node] = true

		for _, next := range graph[node] {
			dfs(next)
		}
	}

	dfs(k)

	for _, edge := range invocations {
		a, b := edge[0], edge[1]

		if !suspicious[a] && suspicious[b] {
			ans := make([]int, n)

			for i := 0; i < n; i++ {
				ans[i] = i
			}

			return ans
		}
	}

	ans := []int{}

	for i := 0; i < n; i++ {
		if !suspicious[i] {
			ans = append(ans, i)
		}
	}

	return ans
}