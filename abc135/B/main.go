package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i])
	}

	var cnt int
	for i := 0; i < n; i++ {
		if p[i] != i+1 {
			cnt++
		}
		if cnt > 2 {
			fmt.Println("NO")
			break
		}
	}
	if cnt <= 2 {
		fmt.Println("YES")
	}
}
