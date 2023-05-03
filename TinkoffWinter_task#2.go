package main

import "fmt"

func main() {
	var a, b, c int
	var x, y, z int
	var n int
	fmt.Scanln(&a, &b, &c)
	fmt.Scanln(&x, &y, &z)
	for i := -(x / a); i <= (y/b + (z / c)); i++ {
		n += ((z / c) - i) + y/b + 1
	}
	fmt.Println(n)
}
