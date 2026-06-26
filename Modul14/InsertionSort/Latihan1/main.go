package main

import "fmt"

func sort(arr *[]int) {
	pArr := *arr
	for i := 1; i < len(pArr); i++ {
		for j := i; j > 0; j-- {
			if pArr[j] < pArr[j-1] {
				temp := pArr[j]
				pArr[j] = pArr[j-1]
				pArr[j-1] = temp
			}
		}
	}
}

func main() {
	var isSorted bool = true
	var arr = []int{}
	for {
		var num int
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		arr = append(arr, num)
	}
	if len(arr) <= 1 {
		fmt.Println("Jumlah data minimal harus 2")
		return
	}
	sort(&arr)
	gap := 0
	if len(arr) > 1 {
		gap = arr[1] - arr[0]
		for i := 0; i < len(arr)-1; i++ {
			if arr[i]+gap != arr[i+1] {
				isSorted = false
				break
			}
		}
		for _, v := range arr {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	} else {
		isSorted = false
	}
	if isSorted {
		fmt.Printf("Data berjarak %d\n", gap)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
