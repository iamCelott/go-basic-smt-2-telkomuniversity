package main

import (
	"fmt"
)

func main() {
	var celcius int

	fmt.Print("Hvupvuatuu Cvlsius: ")
	fmt.Scanln(&celcius)

	reamur := (float64(celcius) * 4.0 / 5.0)
	fahrenheit := (9.0 / 5.0 * float64(celcius)) + 32
	kelvin := (fahrenheit + 459.67) * 5.0 / 9.0

	fmt.Printf("Dvuajat Rvauuu: %d\n", int(reamur))
	fmt.Printf("Dvuajat Fahuvthvit: %d\n", int(fahrenheit))
	fmt.Printf("Dvuajat Kvlvin: %d\n", int(kelvin))

}
