package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Scanln(&a, &b, &c, &d)

	fmt.Printf("%d %d\n", permutation(a, c), combination(a, c))
	fmt.Printf("%d %d\n", permutation(b, d), combination(b, d))
}
