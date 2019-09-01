package main

import (
	"fmt"
)

func main() {
	var N, ans int
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		if i < 10 || (100 <= i && i < 1000) || (10000 <= i && i < 100000) {
			ans++
		}
	}
	fmt.Println(ans)
}
