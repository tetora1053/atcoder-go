package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, first, second int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)

	a := make([]int, n)
	for i := range a {
		sc.Scan()
		a[i], _ = strconv.Atoi(sc.Text())

		if i == 1 {
			if a[0] >= a[1] {
				first = a[0]
				second = a[1]
			} else {
				first = a[1]
				second = a[0]
			}
		}
		if i > 1 {
			if a[i] > first {
				second = first
				first = a[i]
			} else if a[i] > second {
				second = a[i]
			}
		}
	}

	if first == second {
		for i := 0; i < n; i++ {
			fmt.Println(first)
		}
	} else {
		for i := range a {
			if a[i] != first {
				fmt.Println(first)
			} else {
				fmt.Println(second)
			}
		}
	}
}
