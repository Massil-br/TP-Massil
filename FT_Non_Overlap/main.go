package main

import "fmt"

func Ft_non_overlap(intervals [][]int) int {

}

func main() {
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}}))
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))
}
