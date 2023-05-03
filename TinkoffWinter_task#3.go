package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var a, b, div int
	for i := 2; i <= n; i++ {
		if n%i == 0 {
			div = i
			break
		}
	}
	b = n - (n / div)
	a = n - b
	fmt.Println(a, b)
}
