package main

import (
	"fmt"
	"strings"
)

func main() {
	var N, K int
	var S string

	fmt.Scanln(&N, &K)
	fmt.Scanln(&S)

	fmt.Println(S[0:K-1] + strings.ToLower(string(S[K-1])) + S[K:])
}
