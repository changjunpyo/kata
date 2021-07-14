package kata

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func backtrack(dp [][]int, s1, s2 string, i, j int) string {
	if i == 0 || j == 0 {
		return ""
	}
	if s1[i-1] == s2[j-1] {
		return backtrack(dp, s1, s2, i-1, j-1) + string(s1[i-1])
	}
	if dp[i][j-1] > dp[i-1][j] {
		return backtrack(dp, s1, s2, i, j-1)
	}
	return backtrack(dp, s1, s2, i-1, j)
}

// LCS returns the longest subsequence common to s1 and s2
func LCS(s1, s2 string) string {
	n, m := len(s1), len(s2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return backtrack(dp, s1, s2, n, m)
}
