package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	mp := make(map[string]int)

	for i := 0; i < 4; i++ {
		mp[s[i:i+1]]++
	}

	if len(mp) == 2 && mp[s[0:1]] == 2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
