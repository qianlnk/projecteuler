/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package main

import (
	"fmt"

	"github.com/qianlnk/projecteuler/lnkmath"
)

func main() {
	pns := lnkmath.LessNPrimeNumber(2000000)
	var res int64 = 0
	for _, pn := range pns {
		res += pn
	}
	fmt.Println(res)
}
