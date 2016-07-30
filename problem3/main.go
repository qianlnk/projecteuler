/*
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/

package main

import (
	"fmt"

	"github.com/qianlnk/projecteuler/lnkmath"
)

func main() {
	res := lnkmath.PrimeFactor(600851475143)
	fmt.Println(res[len(res)-1])
}
