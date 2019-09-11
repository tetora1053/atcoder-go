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

	a := make([]int, N)
	b := make([]int, N)
	c := make([]int, N-1)

	for i := range a {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())
	}

	for i := range b {
		sc.Scan()
		b[i], _ = strconv.Atoi(sc.Text())
	}

	for i := range c {
		sc.Scan()
		c[i], _ = strconv.Atoi(sc.Text())
	}

	ans := b[0]
	pre := a[0]

	for i := 1; i < N; i++ {
		ans += b[i]
		if pre+1 == a[i] {
			ans += c[pre-1]
		}
		pre = a[i]
	}
	fmt.Println(ans)
}
