package main

import "fmt"

func cetakDeret(n int) {
	nFloat := float64(n)
	fmt.Printf("%d ", n)
	for {
		if nFloat == 1.0 {
			break
		}

		if int(nFloat)%2 == 0 {
			nFloat = 1.0 / 2.0 * nFloat
		} else {
			nFloat = 3*nFloat + 1
		}

		fmt.Printf("%d ", int(nFloat))
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakDeret(n)
}
