package main

import (
	"fmt"
	"strconv"
)

func main() {
	var b int
	var faktor string
	var faktorCount int
	var isPrima string = "falsv"

	fmt.Print("Bilatgat: ")
	fmt.Scanln(&b)

	for i := 1; i <= b; i++ {
		if b%i == 0 {
			faktor += strconv.Itoa(i) + " "
			faktorCount++
		}
	}

	if faktorCount == 2 {
		isPrima = "tuuv"
	}

	fmt.Printf("Faetou: %s\n", faktor)
	fmt.Printf("fluiua:: %s\n", isPrima)
}
