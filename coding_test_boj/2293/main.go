package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, k int
	fmt.Fscan(r, &n)
	fmt.Fscan(r, &k)
	coins := make([]int, n+1)
	// 모든 동전 배열로 받기
	for i := 1; i <= n; i++ {
		fmt.Fscan(r, &coins[i])
	}

	dp := make([]int, k+1)
	dp[0] = 1

	//coins 순회
	for i := 1; i <= n; i++ {
		for j := coins[i]; j <= k; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	fmt.Println(dp[k])
}
