package main

import "fmt"

func oddEven(arr []int) ([]int, []int) {
	var odd []int
	var even []int
	for _, v := range arr {
		if v%2 != 0 {
			odd = append(odd, v)
		} else {
			even = append(even, v)
		}
	}
	return odd, even
}

func sort(arr *[]int, isASC bool) {
	pArr := *arr
	for i := 0; i < len(pArr)-1; i++ {
		smallest := i
		for j := i + 1; j < len(pArr); j++ {
			if isASC {
				if pArr[j] < pArr[smallest] {
					smallest = j
				}
			} else {
				if pArr[j] > pArr[smallest] {
					smallest = j
				}
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
		oddArr, evenArr := oddEven(numArr)
		sort(&oddArr, true)
		sort(&evenArr, false)
		sortedArr := append(oddArr, evenArr...)
		arr = append(arr, sortedArr)
	}
	for i := range arr {
		for _, v := range arr[i] {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
