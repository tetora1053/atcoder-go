package main

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
	k, _ := strconv.Atoi(sl[0])
	x, _ := strconv.Atoi(sl[1])

	min := x - k + 1
	if min < -1000000 {
		min = -1000000
	}
	max := x + k - 1
	if 1000000 < max {
		max = 1000000
	}

	for i := min; i <= max; i++ {
		fmt.Print(i, " ")
	}
}
