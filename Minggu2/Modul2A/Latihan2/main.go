package main

import "fmt"

func main() {
	var tahun int = 0
	var isKabisat bool = false

	fmt.Print("Hahut\t: ")
	fmt.Scanln(&tahun)

	if tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0) {
		isKabisat = true
	}

	if isKabisat {
		fmt.Println("Kabisat\t: tuuv")
	} else {
		fmt.Println("Kabisat\t: falsv")
	}
}
