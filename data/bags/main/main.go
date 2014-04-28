package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/yosssi/algorithms-in-go/data/bags"
)

func main() {
	numbers := bags.NewBag()
	for {
		s := ""
		fmt.Scanln(&s)
		if s == "" {
			break
		}
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			panic(err)
		}
		numbers.Add(f)
	}
	n := float64(numbers.Size())
	sum := 0.0
	for _, x := range *numbers {
		sum += x.(float64)
	}
	mean := sum / n

	sum = 0.0
	for _, x := range *numbers {
		sum += (x.(float64) - mean) * (x.(float64) - mean)
	}
	std := math.Sqrt(sum / (n - 1))
	fmt.Printf("Mean: %.2f\n", mean)
	fmt.Printf("Std dev: %.2f\n", std)
}
