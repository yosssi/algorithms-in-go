package main

import (
	"fmt"

	"github.com/yosssi/algorithms-in-go/data/stacks"
)

func main() {
	stack := stacks.NewStack()

	for {
		s := ""
		fmt.Scanf("%v\n", &s)
		if s == "" {
			break
		}
		stack.Push(s)
	}

	for !stack.IsEmpty() {
		fmt.Printf("%+v\n", stack.Pop())
	}
}
