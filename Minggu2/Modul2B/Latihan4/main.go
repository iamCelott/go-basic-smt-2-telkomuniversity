package main

import "fmt"

func main() {
	var k int
	var hasil float64 = 1

	fmt.Print("Nilai K = ")
	fmt.Scanln(&k)

	for i := 0; i <= k; i++ {
		fk := float64((4*i+2)*(4*i+2)) / float64((4*i+1)*(4*i+3))
		hasil *= fk
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", hasil)
}
