/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

//conditions:
//a < b < c
//a^2 + b^2 = c^2
//a + b + c=1000
//a + b > c
//result:
//250 < b < c < 500

package main

import (
	"fmt"
)

func main() {
	n := 1000
	for b := n / 4; b < n/2; b++ {
		for c := n / 4; c < n/2; c++ {
			a := n - b - c
			if a*a+b*b == c*c {
				fmt.Println(a*b*c, "=", a, "*", b, "*", c)
			}
		}
	}
}
