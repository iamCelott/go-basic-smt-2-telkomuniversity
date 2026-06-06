package main

import (
	"fmt"
	"math"
)

func main() {
	var jumlah float64
	var i int

	for i = 1; ; i++ {

		si := 1.0 / float64(2*i-1)
		if i%2 == 0 {
			si = -si
		}

		jumlah += si

		sip1 := 1.0 / float64(2*(i+1)-1)
		if (i+1)%2 == 0 {
			sip1 = -sip1
		}

		pi1 := 4 * jumlah
		pi2 := 4 * (jumlah + sip1)

		if math.Abs(pi1-pi2) <= 0.00001 {
			fmt.Println("N suku pertama: 1000000")
			fmt.Printf("Hasil PI: %.10f\n", pi1)
			fmt.Printf("Hasil PI: %.10f\n", pi2)
			fmt.Println("Pada i ke:", i)
			break
		}
	}
}
