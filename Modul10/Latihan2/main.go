package main

import "fmt"

const NMAX = 1000

func main() {
	var x, y int
	var ikan [NMAX]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	jumlahWadah := (x + y - 1) / y
	var totalSemua float64

	for w := 0; w < jumlahWadah; w++ {
		var totalWadah float64

		awal := w * y
		akhir := awal + y
		if akhir > x {
			akhir = x
		}

		for i := awal; i < akhir; i++ {
			totalWadah += ikan[i]
		}

		totalSemua += totalWadah

		fmt.Printf("%.2f ", totalWadah)
	}
	fmt.Println()

	rata := totalSemua / float64(jumlahWadah)
	fmt.Printf("%.2f\n", rata)
}