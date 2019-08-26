package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Task struct {
	a, b int
}

type Tasks []Task

func (t Tasks) Len() int { return len(t) }
func (t Tasks) Less(i, j int) bool {
	return t[i].b < t[j].b
}
func (t Tasks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func main() {
	var n, sum int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	tasks := make([]Task, n)
	for i := range tasks {
		sc.Scan()
		tasks[i].a, _ = strconv.Atoi(sc.Text())
		sc.Scan()
		tasks[i].b, _ = strconv.Atoi(sc.Text())
	}
	sort.Sort(Tasks(tasks))

	for _, task := range tasks {
		sum += task.a
		if sum > task.b {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
