/*
A palindromic number reads the same both ways. The largest palindrome made
from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
*/
package main

import (
	"fmt"

	"github.com/qianlnk/projecteuler/lnkmath"
)

func main() {
	var left, right int64
	var largest int64
	for left = 999; left > 99; left-- {
		for right = 999; right > 99; right-- {
			pal := left * right
			if lnkmath.IsPalindromic(pal) {
				fmt.Println(pal, "=", left, "*", right)
				if largest < pal {
					largest = pal
				}
			}
		}
	}
	fmt.Println(largest)
}
