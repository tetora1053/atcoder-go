package main

import (
	"fmt"
)

func main() {
	var w, h, x, y float64
	var halfArea float64
	fmt.Scan(&w, &h, &x, &y)

	halfArea = w * h / 2

	midX := w / 2
	midY := h / 2

	if midX == x && midY == y {
		fmt.Println(halfArea, 1)
	} else {
		fmt.Println(halfArea, 0)
	}
}
