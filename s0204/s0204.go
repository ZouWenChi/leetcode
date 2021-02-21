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
	count := 0
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return count
}

func main() {
	fmt.Println(countPrimes(10))
}
