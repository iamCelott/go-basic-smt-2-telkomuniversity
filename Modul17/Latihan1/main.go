package main

import "fmt"

func main() {
	var x, jumlah float64
	var cacah int

	fmt.Scan(&x)

	for x != 9999 {
		jumlah += x
		cacah++
		fmt.Scan(&x)
	}

	if cacah > 0 {
		fmt.Println("Rerata =", jumlah/float64(cacah))
	} else {
		fmt.Println("Tidak ada data")
	}
}
