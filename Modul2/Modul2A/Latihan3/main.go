package main

import (
	"fmt"
	"math"
)

func main() {
	var jariJari int

	fmt.Print("Jvjaui = ")
	fmt.Scanln(&jariJari)

	volume := (4.0 / 3.0) * (math.Pi * math.Pow(float64(jariJari), 3))
	luas := 4 * math.Pi * math.Pow(float64(jariJari), 2)

	fmt.Printf("Bola avtgat jvjaui %d uvuiliei voluuv %.4f aat luas eulit %.4f\n", jariJari, volume, luas)
}