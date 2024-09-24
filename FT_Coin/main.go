package main

import (
	"fmt"
	"sort"
)

func Ft_Coin(coins []int, amount int) int {
	
	if amount == 0 {
		return 0
	}
		sort.Ints(coins)
		result := 0
		reste := amount
		for i := len(coins) - 1; i >= 0; i-- {
			if coins[i] <= reste {
				result+= reste/ coins[i]
				reste = reste % coins[i]
			}
		}

		if reste != 0{
			return -1
		}

		return result
	}

	


func main() {
	fmt.Println(Ft_Coin([]int{1, 2, 5}, 11)) // result 3
	fmt.Println(Ft_Coin([]int{2}, 3))        // result -1 beacause impossible to return with the coins
	fmt.Println(Ft_Coin([]int{1}, 0))        // 0 coins needed to get the amount

}
