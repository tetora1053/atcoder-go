package main

import (
	"fmt"
	"strconv"
)

func main() {
	var S string
	fmt.Scanln(&S)

	former, _ := strconv.Atoi(S[0:2])
	latter, _ := strconv.Atoi(S[2:])

	if 1 <= former && former <= 12 && 1 <= latter && latter <= 12 {
		fmt.Println("AMBIGUOUS")
	} else if 1 <= former && former <= 12 {
		fmt.Println("MMYY")
	} else if 1 <= latter && latter <= 12 {
		fmt.Println("YYMM")
	} else {
		fmt.Println("NA")
	}
}
