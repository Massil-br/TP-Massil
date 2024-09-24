package main

import (
	"fmt"
)

func Ft_Profit(prices []int) int {
	min := prices[0]

	profitmax := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			profit := prices[i] - min

			if profit > profitmax {
				profitmax = profit
			}
		}

	}
	return profitmax
}

func main() {
	fmt.Println(Ft_Profit([]int{7, 1, 5, 3, 6, 4}))

	fmt.Println(Ft_Profit([]int{7, 6, 4, 3, 1}))
}
