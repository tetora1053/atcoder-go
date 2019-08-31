package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Restraunt struct {
	n int
	s string
	p int
}

type Restraunts []Restraunt

func (rs Restraunts) Len() int {
	return len(rs)
}

func (rs Restraunts) Less(i, j int) bool {
	if rs[i].s != rs[j].s {
		return rs[i].s < rs[j].s
	}
	return rs[i].p > rs[j].p
}

func (rs Restraunts) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func main() {
	var N int
	fmt.Scan(&N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	var rs Restraunts = make([]Restraunt, N)
	for i := range rs {
		sc.Scan()
		rs[i].n, rs[i].s = i+1, sc.Text()
		sc.Scan()
		rs[i].p, _ = strconv.Atoi(sc.Text())
	}
	sort.Sort(rs)
	for i := range rs {
		fmt.Println(rs[i].n)
	}
}
