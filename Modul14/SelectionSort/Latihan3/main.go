package main

import "fmt"

func findMedian(arr []int) int {
	median := 0
	lenArr := len(arr)
	if lenArr%2 == 0 {
		median = (arr[lenArr/2] + arr[lenArr/2-1]) / 2
	} else {
		median = arr[lenArr/2]
	}
	return median
}

func sort(arr *[]int) {
	pArr := *arr
	for i := range pArr {
		var smallest int = i
		for j := i + 1; j < len(pArr); j++ {
			if pArr[j] < pArr[smallest] {
				smallest = j
			}
		}
		prev := pArr[i]
		pArr[i] = pArr[smallest]
		pArr[smallest] = prev
	}
}

func main() {
	var arr = [][]int{}
	var subArray = []int{}
	var medians = []int{}
	for {
		var num int
		fmt.Scan(&num)

		if num == -5313 {
			break
		}

		if num == 0 {
			sort(&subArray)
			median := findMedian(subArray)
			medians = append(medians, median)
			arr = append(arr, subArray)
			continue
		}
		subArray = append(subArray, num)
	}
	for _, v := range medians {
		fmt.Println(v)
	}
}
