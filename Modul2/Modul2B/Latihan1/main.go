package main

import "fmt"

func main() {
	var merah, kuning, hijau, ungu string
	var isSuccess bool = false
	var passedCount int = 0

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scanln(&merah, &kuning, &hijau, &ungu)

		if merah == "merah" && kuning == "kuning" && hijau == "hijau" && ungu == "ungu" {
			passedCount++
		}
	}

	if passedCount == 5 {
		isSuccess = true
	}

	fmt.Printf("BERHASIL: %t\n", isSuccess)
}