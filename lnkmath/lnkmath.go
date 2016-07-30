package lnkmath

import (
//"fmt"
)

//判断是否为回文数
//return true if the number is palindromic number, or false if the number is not.
func IsPalindromic(num int64) bool {
	tmp := num
	var n int64

	for tmp != 0 {
		n = n*10 + tmp%10
		tmp = tmp / 10
	}

	return num == n
}

//最大公约数
//reuturn the greatest common divisor of a and b
func GCD(a, b int64) int64 {
	if a > b {
		return GCD(b, a)
	}

	if a == 0 {
		return b
	}

	return GCD(b%a, a)
}

//最小公倍数
//return the lowest common multiple of a and b
func LCM(a, b int64) int64 {
	return a * b / GCD(a, b)
}

//求质因数
//return prime factor list
func PrimeFactor(num int64) []int64 {
	var i int64
	res := make([]int64, 0)
	tmpnum := num
	for i = 2; i < num/2; i++ {
		if tmpnum%i == 0 {
			tmpnum /= i
			res = append(res, i)
		}
	}

	return res
}
