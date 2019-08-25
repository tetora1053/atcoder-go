package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var n, d, ans int
	fmt.Scan(&n, &d)

	xS := make([][]float64, n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i := range xS {
		xS[i] = make([]float64, d)
		for j := range xS[i] {
			sc.Scan()
			tmp, _ := strconv.Atoi(sc.Text())
			xS[i][j] = float64(tmp)
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			var sum float64
			for k := 0; k < d; k++ {
				sum += (xS[i][k] - xS[j][k]) * (xS[i][k] - xS[j][k])
			}
			res := math.Sqrt(sum)
			if res == float64(int64(res)) {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
