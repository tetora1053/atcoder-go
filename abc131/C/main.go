package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	x := (b / c) + (b / d) - (b / lcm(c, d))
	y := ((a - 1) / c) + ((a - 1) / d) - ((a - 1) / lcm(c, d))

	fmt.Println(b - a + 1 - (x - y))
}

func gcd(a, b int) int {
	for {
		if a%b == 0 {
			return b
		}
		tmp := b
		b = a % b
		a = tmp
	}
	return b
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
