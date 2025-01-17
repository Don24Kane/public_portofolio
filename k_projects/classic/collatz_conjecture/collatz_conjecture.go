# by Don24Kane

# GOlang

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("enter your number: ")
	fmt.Scan(&n)

	steps := 0

	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		steps++
	}

	fmt.Printf("the number of steps it took for n to reach the value of 1 is: %d\n", steps)
}
