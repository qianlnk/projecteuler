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
