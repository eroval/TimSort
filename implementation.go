package main

import (
	"fmt"
	"sort"
)

const (
	minRun = 32
)

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}

func merge(arr []int, tmp []int, left, middle, right int) {
	i, j, k := left, middle+1, left
	for i <= middle && j <= right {
		if arr[i] < arr[j] {
			tmp[k] = arr[i]
			i++
		} else {
			tmp[k] = arr[j]
			j++
		}
		k++
	}
	for i <= middle {
		tmp[k] = arr[i]
		k++
		i++
	}
	for j <= right {
		tmp[k] = arr[j]
		k++
		j++
	}
	for i := left; i <= right; i++ {
		arr[i] = tmp[i]
	}
}

func timSort(arr []int) {
	n := len(arr)
	tmp := make([]int, n)
	for i := 0; i < n; i += minRun {
		left := i
		right := min(i+minRun-1, n-1)
		insertionSort(arr[left : right+1])
	}
	for size := minRun; size < n; size *= 2 {
		for left := 0; left < n; left += 2 * size {
			middle := left + size - 1
			right := min(left+2*size-1, n-1)
			merge(arr, tmp, left, middle, right)
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	arr := []int{5, 1, 4, 2, 8}
	timsort(arr)
	fmt.Println(arr) // [1, 2, 4, 5, 8]
}
