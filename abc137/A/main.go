package main

/**
 * -100 <= A,B int <= 100
 * A + B, A - B, A * Bの中で最大の数を出力
 */

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()
	sl := strings.Split(s, " ")
	a, _ := strconv.Atoi(sl[0])
	b, _ := strconv.Atoi(sl[1])

	calcs := []int{
		a + b,
		a - b,
		a * b,
	}

	var max int

	for i, val := range calcs {
		if i == 0 {
			max = val
			continue
		}
		if val > max {
			max = val
		}
	}
	fmt.Println(max)
}
