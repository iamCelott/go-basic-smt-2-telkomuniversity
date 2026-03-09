package main

import "fmt"

func main() {
	var n int
	var pita string
	var count int = 0

	fmt.Print("N: ")
	fmt.Scanln(&n)

	if n != 0 {
		for i := 1; i <= n; i++ {
			var bunga string
			fmt.Printf("Bunga %d: ", i)
			fmt.Scanln(&bunga)

			if bunga == "SELESAI" {
				break
			}

			if i > 1 {
				pita += " " + "-" + " " + bunga
			} else {
				pita += bunga
			}

			count++
		}
	}

	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Bunga: %d\n", count)
}