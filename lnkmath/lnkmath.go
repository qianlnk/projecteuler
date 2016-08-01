package lnkmath

import (
	"math"
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

//等差数列和
//return sum of arithmetic progression
func SOAP(a1, d float64, n int64) float64 {
	return a1*float64(n) + (float64(n)*(float64(n)-1)/2)*d
}

//等比数列和
//return sum of geometric progression
func SOGP(a1, q float64, n int64) float64 {
	if q == 1 {
		return a1 * float64(n)
	} else {
		return a1 * ((1 - math.Pow(q, float64(n))) / (1 - q))
	}
}

//1的平方加到n的平方
//The sum of the squares of the natural numbers
func SOSONN(n int64) int64 { //sum of squares of natural numbers
	return n * (n + 1) * (2*n + 1) / 6
}

//根據埃拉托斯特尼筛法求前n個質數
//return first n prime number
func FirstNPrimeNumber(n int64) []int64 {
	if n == 1 {
		return []int64{2}
	}
	smallPnLst := []int{2, 3, 5, 7}
	step := 100
	nnlst := make([]bool, step)
	var pnlst []int64
	for i := 2; i < step; i++ {
		nnlst[i] = true
	}

	for _, p := range smallPnLst {
		for j := 2; j < step; j++ {
			if p*j < step {
				nnlst[p*j] = false
			}
		}
	}

	for i, v := range nnlst {
		if v {
			pnlst = append(pnlst, int64(i))
		}
	}

	fornum := 1
	for int64(len(pnlst)) < n {
		for i := fornum * step; i < step*(fornum+1); i++ {
			nnlst = append(nnlst, true)
		}

		for _, p := range pnlst {
			for j := 2; j < step*(fornum+1); j++ {
				if int(p)*j < step*fornum {
					continue
				}
				if int(p)*j < step*(fornum+1) {
					nnlst[int(p)*j] = false
				} else {
					break
				}
			}
		}

		for i, v := range nnlst[(step * fornum):(step * (fornum + 1))] {
			if v {
				pnlst = append(pnlst, int64(i+step*fornum))
			}
		}

		fornum++
	}

	return pnlst[:n]
}

//根據埃拉托斯特尼筛法求小於n的質數
func LessNPrimeNumber(n int64) []int64 {
	if n <= 2 {
		return nil
	}
	smallPnLst := []int{2, 3, 5, 7}
	step := 100
	nnlst := make([]bool, step)
	var pnlst []int64
	for i := 2; i < step; i++ {
		nnlst[i] = true
	}

	for _, p := range smallPnLst {
		for j := 2; j < step; j++ {
			if p*j < step {
				nnlst[p*j] = false
			}
		}
	}

	for i, v := range nnlst {
		if v && int64(i) < n {
			pnlst = append(pnlst, int64(i))
		}
	}

	fornum := 1
	for {
		for i := fornum * step; i < step*(fornum+1); i++ {
			nnlst = append(nnlst, true)
		}

		for _, p := range pnlst {
			for j := 2; j < step*(fornum+1); j++ {
				if int(p)*j < step*fornum {
					continue
				}
				if int(p)*j < step*(fornum+1) {
					nnlst[int(p)*j] = false
				} else {
					break
				}
			}
		}
		out := false
		for i, v := range nnlst[(step * fornum):(step * (fornum + 1))] {
			if int64(i+step*fornum) >= n {
				out = true
				break
			}

			if v && int64(i+step*fornum) < n {
				pnlst = append(pnlst, int64(i+step*fornum))
			}
		}

		if out {
			break
		}

		fornum++
	}
	return pnlst
}
