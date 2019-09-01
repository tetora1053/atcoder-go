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

	var cur int
	for i := 0; i < N; i++ {
		sc.Scan()
		h, _ := strconv.Atoi(sc.Text())

		if h < cur-1 {
			fmt.Println("No")
			return
		} else if h > cur {
			cur = h
		}
	}
	fmt.Println("Yes")
}
