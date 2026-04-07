package main

import "fmt"

func barisan(i, n int) {
	if i > n {
		return
	}
	if i%2 != 0 {
		fmt.Print(i, " ")
	}
	barisan(i+1, n)
}

func main() {
	var n int
	fmt.Scan(&n)
	barisan(1, n)
	fmt.Println()
}
