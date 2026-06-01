package main

import "fmt"

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func fogoh(a int) int {
	return f(g(h(a)))
}

func gohof(b int) int {
	return g(h(f(b)))
}

func hofog(c int) int {
	return h(f(g(c)))
}

func main() {
	var a, b, c int

	fmt.Scanln(&a, &b, &c)

	fmt.Println(fogoh(a))
	fmt.Println(gohof(b))
	fmt.Println(hofog(c))
}
