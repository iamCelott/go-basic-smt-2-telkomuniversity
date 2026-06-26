// Dharma Chandra Viriya
// S1IF-13-02
// 109082500052
package main

import "fmt"

func main() {
	var arr [21]int
	totalVote := 0
	validVote := 0
	for {
		var suara int
		fmt.Scan(&suara)

		if suara == 0 {
			break
		}
		totalVote++
		if suara >= 1 && suara <= 20 {
			validVote++
			arr[suara]++
		}
	}
	ketua := 1
	for i := 2; i <= 20; i++ {
		if arr[i] > arr[ketua] {
			ketua = i
		}
	}
	wakil := 0
	for i := 1; i <= 20; i++ {
		if i == ketua {
			continue
		}
		if wakil == 0 || arr[i] > arr[wakil] {
			wakil = i
		}
	}
	fmt.Println("Suara masuk:", totalVote)
	fmt.Println("Suara sah:", validVote)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil Ketua:", wakil)
}
