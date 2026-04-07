package main

import "fmt"

func rekursif(i int, a int, b int) {
	if i == 0 {
		return
	}
	fmt.Print(a, " ")
	rekursif(i-1, b, a+b)
}

func main() {
	rekursif(11, 0, 1)
	fmt.Println()
}
