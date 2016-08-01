package lnkmath

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	fmt.Println(GCD(20, 8))
}

func TestLCM(t *testing.T) {
	fmt.Println(LCM(6, 20))
}

func TestIsPalindromic(t *testing.T) {
	fmt.Println(IsPalindromic(123321))
	fmt.Println(IsPalindromic(123432))
}

func TestPrimeFactor(t *testing.T) {
	res := PrimeFactor(13195)
	for _, v := range res {
		fmt.Println(v)
	}
}

func TestSOAP(t *testing.T) {
	fmt.Println(SOAP(1, 1, 5))
}

func TestSOGP(t *testing.T) {
	fmt.Println(SOGP(2, 1, 5))
	fmt.Println(SOGP(1, 2, 5))
}

func TestPrimeNumber(t *testing.T) {
	fmt.Println(FirstNPrimeNumber(10001))
}
