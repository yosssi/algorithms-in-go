package main

import (
	"fmt"
	"strconv"

	"github.com/yosssi/algorithms-in-go/data/fifoqueues"
)

func main() {
	q := fifoqueues.NewQueue()
	for {
		s := ""
		fmt.Scanln(&s)
		if s == "" {
			break
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		q.Enqueue(i)
	}
	n := q.Size()
	a := make([]int, n)
	for idx, _ := range a {
		a[idx] = q.Dequeue().(int)
	}
	for _, v := range a {
		fmt.Println(v)
	}
}
