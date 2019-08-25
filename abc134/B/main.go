package main

import (
	"fmt"
)

func main() {
	var n, d int
	fmt.Scan(&n, &d)

	treePerWatcher := 2*d + 1
	if n%treePerWatcher == 0 {
		fmt.Println(n / treePerWatcher)
	} else {
		fmt.Println(n/treePerWatcher + 1)
	}
}
