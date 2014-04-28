package main

import (
	"fmt"

	"github.com/yosssi/algorithms-in-go/data/bags"
)

func main() {
	numbers := bags.NewBag()
	fmt.Println(numbers.Size())
}
