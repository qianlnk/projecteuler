/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import (
	"fmt"

	"github.com/qianlnk/projecteuler/lnkmath"
)

func main() {
	var res, i int64
	res = 1
	for i = 2; i <= 20; i++ {
		res = lnkmath.LCM(res, i)
	}
	fmt.Println(res)
}
