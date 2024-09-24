package main

import (
	"fmt"
)

func Ft_Missing(nums []int) int {
	number := len(nums)
	expect := number * (number + 1) / 2

	real := 0

	for _, num := range nums {
		real += num
	}

	return expect - real
}

func main() {
	fmt.Println(Ft_Missing([]int{3, 1, 2}))
	fmt.Println(Ft_Missing([]int{0, 1}))
	fmt.Println(Ft_Missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
