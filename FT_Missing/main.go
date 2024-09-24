package main

import (
	"fmt"
	"sort"
)

func Ft_Missing(nums []int) int {
	sort.Ints(nums)
	var res int
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			res = i
		}
	}
	return res
}

func main() {
	fmt.Println(Ft_Missing([]int{3, 1, 2}))
	fmt.Println(Ft_Missing([]int{0, 1}))
	fmt.Println(Ft_Missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
