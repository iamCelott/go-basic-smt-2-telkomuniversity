// package main
// import “fmt”
// func main() {
// var nam float64
// var nmk string
// fmt.Print(“Nilai akhir mata kuliah: “)
// fmt.Scanln(&nam)
// if nam > 80 {
// nam = “A”
// }
// if nam > 72.5 {
// nam = “AB”
// }
// if nam > 65 {
// nam = “B”
// }
// if nam > 57.5 {
// nam = “BC”
// }
// if nam > 50 {
// nam = “C”
// }
// if nam > 40 {
// nam = “D”
// } else if nam <= 40 {
// nam = “E”
// }
// fmt.Println(“Nilai mata kuliah: “, nmk)}

package main

import "fmt"

func main() {
	var nam float64
	var nmk string
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)
	if nam > 80 && nam <= 100 {
		nmk = "A"
	}
	if nam > 72.5 && nam <= 80 {
		nmk = "AB"
	}
	if nam > 65 && nam <= 72.5 {
		nmk = "B"
	}
	if nam > 57.5 && nam <= 65 {
		nmk = "BC"
	}
	if nam > 50 && nam <= 57.5 {
		nmk = "C"
	}
	if nam > 40 && nam <= 50 {
		nmk = "D"
	}
	if nam <= 40 {
		nmk = "E"
	}
	fmt.Println("Nilai mata kuliah: ", nmk)
}
