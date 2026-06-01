// Dharma Chandra Viriya
// S1IF-13-02
// 109082500052
package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func isiArray(n int) {
	/* I.S. terdefinisi integer n, dan sejumlah n data
	   sudah siap pada piranti masukan.
	   F.S. array data berisi n bilangan yang dibaca
	   dari piranti masukan. */

	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	/* mengembalikan posisi k dalam array data dengan n elemen.
	   Posisi dimulai dari 0. Jika tidak ada, kembalikan -1. */

	left := 0
	right := n - 1

	for left <= right {
		mid := (left + right) / 2

		if data[mid] == k {
			return mid
		} else if data[mid] < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func main() {
	/* I.S. terdefinisi bilangan bulat positif n dan k
	   pada piranti masukan, serta n buah data integer
	   positif yang sudah terurut membesar.
	   F.S. tercetak posisi data k dalam array jika ada,
	   atau "TIDAK ADA" jika k tidak ditemukan. */

	var n, k int

	fmt.Scan(&n, &k)

	isiArray(n)

	idx := posisi(n, k)

	if idx == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(idx)
	}
}