package main

import "fmt"

func faktor(i, n int) {
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	if i == n {
		return
	}
	faktor(i+1, n)
}

func main() {
	var n int
	fmt.Scan(&n)
	
	faktor(1, n)
	fmt.Println()
}
