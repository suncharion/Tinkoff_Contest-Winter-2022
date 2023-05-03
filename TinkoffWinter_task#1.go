package main

import (
	"bufio"
	"fmt"
	"os"
)

func main2() {
	var n, space, countBad int
	var flag bool
	var yb string
	fmt.Scanln(&n)
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	fmt.Scanln(&yb)
	yb += " "
	flag = false
	for i := 0; i < (len(line) - 1); i++ {
		if line[i] == 32 {
			space++
			if flag == true {
				countBad++
			}
			flag = false
		}
		if i == (len(line) - 2) {
			if flag == true {
				countBad++
			}

		} else {
			if (line[i+1]) != 32 && (line[i]) != 32 && (yb[i-space]) == (yb[i+1-space]) {
				flag = true
			}
		}
	}
	fmt.Println(countBad)
}
