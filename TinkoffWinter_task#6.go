package main

import (
	"fmt"
)

func main() {
	var n, max_value, b int
	fmt.Scanln(&n)
	var start_values = make(map[int]bool)
	for k := 0; k < n; k++ {
		fmt.Scanln(&b)
		start_values[b] = true

		for key, _ := range start_values {
			if max_value < (b ^ key) {
				max_value = b ^ key
			}
		}
		fmt.Println(max_value)
	}
}
