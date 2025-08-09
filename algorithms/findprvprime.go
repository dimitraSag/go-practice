package main

import (
	"fmt"
	"math"
)

// Check if a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// FindPrevPrime returns the greatest prime number <= nb or 0 if none
func FindPrevPrime(nb int) int {
	for i := nb; i >= 2; i-- {
		if isPrime(i) {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(FindPrevPrime(5))  // Output: 5
	fmt.Println(FindPrevPrime(4))  // Output: 3
	fmt.Println(FindPrevPrime(1))  // Output: 0
	fmt.Println(FindPrevPrime(2))  // Output: 2
}
