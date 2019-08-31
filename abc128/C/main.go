package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		N : スイッチの個数
		M : 電球の個数
	*/
	var N, M, ans int
	fmt.Scan(&N, &M)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	a := make([]int, N)
	for i := 0; i < M; i++ {
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())

		for j := 0; j < k; j++ {
			sc.Scan()
			s, _ := strconv.Atoi(sc.Text())
			a[s-1] |= 1 << uint(i)
		}
	}

	var p int
	for i := 0; i < M; i++ {
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())
		p |= b << uint(i)
	}

	for sState := 0; sState < (1 << uint(N)); sState++ {
		q := 0
		for i := 0; i < N; i++ {
			if sState>>uint(i)&1 == 1 {
				q ^= a[i]
			}
		}
		if p == q {
			ans++
		}
	}
	fmt.Println(ans)
}
