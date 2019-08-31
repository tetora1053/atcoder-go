package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	d := make([]int, n)
	for i := range d {
		sc.Scan()
		d[i], _ = strconv.Atoi(sc.Text())
	}

	sort.Ints(d)

	fmt.Println(d[len(d)/2] - d[len(d)/2-1])
}
