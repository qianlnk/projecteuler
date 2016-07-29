/*
Each new term in the Fibonacci sequence is generated by adding the previous two terms.
By starting with 1 and 2, the first 10 terms will be:
1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
By considering the terms in the Fibonacci sequence whose values do not exceed four
million, find the sum of the even-valued terms.
*/

package main

import (
	"fmt"
)

func main() {
	var limit int64 = 4000000
	fibs := []int64{1, 2}
	sum := fibs[1]
	for {
		fib := fibs[0] + fibs[1]
		if fib <= limit {
			fibs[0] = fibs[1]
			fibs[1] = fib
			if fib%2 == 0 {
				sum += fib
			}
		} else {
			break
		}
	}
	fmt.Println(sum)
}