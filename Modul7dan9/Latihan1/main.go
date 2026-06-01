package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	pusat  titik
	radius int
}

func jarak(p, q titik) float64 {
	return math.Sqrt(
		math.Pow(float64(p.x-q.x), 2) +
			math.Pow(float64(p.y-q.y), 2),
	)
}

func didalam(c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= float64(c.radius)
}

func main() {
	var c1, c2 lingkaran
	var p titik

	fmt.Scan(&c1.pusat.x, &c1.pusat.y, &c1.radius)

	fmt.Scan(&c2.pusat.x, &c2.pusat.y, &c2.radius)

	fmt.Scan(&p.x, &p.y)

	dalam1 := didalam(c1, p)
	dalam2 := didalam(c2, p)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}