package main

import (
	"fmt"
)

func isPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	fmt.Println(x)
	return true
}

func countPrimes(n int) int {
	r := 0
	for x := 2; x < n; x++ {
		if isPrime(x) {
			r++
		}
	}
	return r
}

func main() {
	fmt.Println(countPrimes(10))
}
