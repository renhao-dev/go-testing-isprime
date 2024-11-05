package main

import "fmt"

func main() {
	n := 2

	_, msg := isPrime(n)

	fmt.Println(msg)
}

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by def", n)
	}
	if n < 0 {
		return false, "Negative numbers are not prime by def"
	}

	for i := 2; i <= n-1; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is divisable by %d", n, i)
		}
	}
	return true, fmt.Sprintf("%d is prime", n)
}
