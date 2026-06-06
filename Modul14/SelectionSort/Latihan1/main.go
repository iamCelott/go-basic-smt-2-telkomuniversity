package main

import "fmt"

func sort(arr *[]int) {
	pArr := *arr
	for i := range pArr {
		var smallest int = i
		for j := i + 1; j < len(pArr); j++ {
			if pArr[j] < pArr[smallest] {
				smallest = j
			}
		}
		pArr[i], pArr[smallest] = pArr[smallest], pArr[i]
	}
}

func main() {
	var arr = [][]int{}
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		var numArr = []int{}
		var nSub int
		fmt.Scan(&nSub)
		for j := 1; j <= nSub; j++ {
			var num int
			fmt.Scan(&num)
			numArr = append(numArr, num)
		}
		sort(&numArr)
		arr = append(arr, numArr)
	}

	for i := range arr {
		for _, v := range arr[i] {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
