package main

import "fmt"

func pangkat(i, x, y int, result *int) {
	if i == y {
		return
	}
	*result *= x
	pangkat(i+1, x, y, result)
}

func main() {
	var x, y int
	var result int = 1
	fmt.Scan(&x, &y)
	pangkat(0, x, y, &result)
	fmt.Println(result)
}
