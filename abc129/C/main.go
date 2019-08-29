package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N, M int
	fmt.Scanln(&N, &M)
	sc := bufio.NewScanner(os.Stdin)
	broken := make(map[int]bool, M)
	for i := 0; i < M; i++ {
		sc.Scan()
		stair, _ := strconv.Atoi(sc.Text())
		broken[stair] = true
	}

	dp := make([]int, N+1)
	dp[0] = 1
	if broken[1] {
		dp[1] = 0
	} else {
		dp[1] = 1
	}

	for i := 2; i <= N; i++ {
		if broken[i] == false {
			dp[i] = dp[i-1] + dp[i-2]
		}
		dp[i] %= 1000000007
	}
	fmt.Println(dp[N])
}
