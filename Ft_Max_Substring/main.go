package main

import "fmt"

func Ft_max_substring(s string) int {
	countmax := 1
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1]+1 {
			count++
		} else {
			if count > countmax {
				countmax = count
			}
			count = 1
		}

	}

	if count > countmax {
		countmax = count
	}

	return countmax
}

func main() {
	fmt.Println(Ft_max_substring("abcabcbb"))
	fmt.Println(Ft_max_substring("bbbbbb"))
}
