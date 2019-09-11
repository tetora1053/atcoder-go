package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N int
	fmt.Scanln(&N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	b := make([]int, N-1)
	for i := range b {
		sc.Scan()
		b[i], _ = strconv.Atoi(sc.Text())
	}

	ans := b[0]
	for i := 0; i < N-2; i++ {
		if b[i] <= b[i+1] {
			ans += b[i]
		} else {
			ans += b[i+1]
		}
	}
	ans += b[N-2]
	fmt.Println(ans)
}
