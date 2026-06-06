package main

import "fmt"

func main() {
	var x, data string
	var n, i int
	var jumlah int = 0
	var posisi int = -1

	fmt.Scan(&x)
	fmt.Scan(&n)

	for i = 1; i <= n; i++ {
		fmt.Scan(&data)

		if data == x {
			jumlah++

			if posisi == -1 {
				posisi = i
			}
		}
	}

	if jumlah > 0 {
		fmt.Println("String ditemukan")
		fmt.Println("Posisi pertama:", posisi)
	} else {
		fmt.Println("String tidak ditemukan")
	}

	fmt.Println("Jumlah kemunculan:", jumlah)

	if jumlah >= 2 {
		fmt.Println("Ada sedikitnya dua kemunculan")
	} else {
		fmt.Println("Tidak ada dua kemunculan")
	}
}
