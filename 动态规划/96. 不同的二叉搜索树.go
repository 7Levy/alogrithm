package main

func numTrees(n int) int {
	dp := make([]int, n+3)
	//动态规划[i] +=动态规划[j-1]*动态规划[i-j]

	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < n+1; i++ {
		for j := i; j >= 1; j-- {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
