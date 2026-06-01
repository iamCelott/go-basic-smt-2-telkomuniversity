package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var element []int

	fmt.Print("Masukkan panjang elemen: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var elem int
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&elem)
		element = append(element, elem)
	}

	fmt.Println("\na. Keseluruhan isi array:")
	for _, v := range element {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Println("\nb. Array dengan indeks ganjil:")
	for i, v := range element {
		if i%2 != 0 {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()

	fmt.Println("\nc. Array dengan indeks genap:")
	for i, v := range element {
		if i%2 == 0 {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()

	var x int
	fmt.Print("\nMasukkan bilangan x: ")
	fmt.Scan(&x)

	fmt.Printf("d. Array dengan indeks kelipatan %d:\n", x)

	if x != 0 {
		for i, v := range element {
			if i%x == 0 {
				fmt.Print(v, " ")
			}
		}
	} else {
		fmt.Println("x tidak boleh 0")
	}
	fmt.Println()

	var deleteIndex int
	fmt.Print("\nMasukkan indeks yang akan dihapus: ")
	fmt.Scan(&deleteIndex)

	element = append(element[:deleteIndex], element[deleteIndex+1:]...)

	fmt.Printf("e. Array setelah indeks %d dihapus:\n", deleteIndex)
	for _, v := range element {
		fmt.Print(v, " ")
	}
	fmt.Println()

	total := 0
	for _, v := range element {
		total += v
	}

	rata := float64(total) / float64(len(element))

	fmt.Println("\nf. Rata-rata:")
	fmt.Println(rata)

	var jumlahKuadrat float64

	for _, v := range element {
		selisih := float64(v) - rata
		jumlahKuadrat += selisih * selisih
	}

	standarDeviasi := math.Sqrt(jumlahKuadrat / float64(len(element)))

	fmt.Println("\ng. Standar deviasi:")
	fmt.Println(standarDeviasi)

	var cari int
	fmt.Print("\nMasukkan bilangan yang dicari frekuensinya: ")
	fmt.Scan(&cari)

	frekuensi := 0
	for _, v := range element {
		if v == cari {
			frekuensi++
		}
	}

	fmt.Printf("h. Frekuensi bilangan %d = %d\n", cari, frekuensi)
}
