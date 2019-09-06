package main

import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scanln(&N, &K)

	var ans, tmp float64

	for i := 1; i <= N; i++ {
		now := i
		tmp = 1.0 / float64(N)
		for now < K {
			now *= 2
			tmp /= 2
		}
		ans += tmp
	}
	fmt.Println(ans)
}
