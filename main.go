package main

import (
	"fmt"

	"github.com/piotrusio/go-pro/algo"
)

func main() {
	nums := []int{2,7,11,15}
	target := 9
	resultA := algo.TwoSumA(nums, target)
	resultB := algo.TwoSumB(nums, target)
	fmt.Println("Total resultA: ", resultA)
	fmt.Println("Total resultB: ", resultB)
}