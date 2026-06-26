// Dharma Chandra Viriya
// S1IF-13-02
// 109082500052
package main

import "fmt"

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var fn, fnr int
	factorial(n, &fn)
	factorial(n-r, &fnr)
	*hasil = fn / fnr
}

func combination(n, r int, hasil *int) {
	var fn, fr, fnr int
	factorial(n, &fn)
	factorial(r, &fr)
	factorial(n-r, &fnr)
	*hasil = fn / (fr * fnr)
}

func main() {
	var a, b, c, d int
	var p1, k1, p2, k2 int
	fmt.Scanln(&a, &b, &c, &d)
	permutation(a, c, &p1)
	combination(a, c, &k1)
	permutation(b, d, &p2)
	combination(b, d, &k2)
	fmt.Printf("%d %d\n", p1, k1)
	fmt.Printf("%d %d\n", p2, k2)
}
