/*
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/

package main

import (
	"fmt"
)

func main() {
	var num int64 = 600851475143
	tmpnum := num
	var i int64
	for i = 2; i < num/2; i++ {
		if tmpnum%i == 0 {
			tmpnum = tmpnum / i
			fmt.Println(i)
		}
	}
}
