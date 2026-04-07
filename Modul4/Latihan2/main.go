package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	for i := 0; i < 8; i++ {
		var waktu int
		fmt.Scan(&waktu)

		if waktu != 301 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var first bool = true
	var winnerName string
	var soalMax, skorMax int

	for {
		var nama string
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		var soal, skor int = 0, 0
		hitungSkor(&soal, &skor)

		if first || soal > soalMax || (soal == soalMax && skor < skorMax) {
			soalMax = soal
			skorMax = skor
			winnerName = nama
			first = false
		}
	}
	fmt.Println(winnerName, soalMax, skorMax)
}
