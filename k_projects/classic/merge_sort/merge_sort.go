# by Don24Kane    

# GOlang

package main

import (
	"fmt"
)

func merge(v []int, left, mid, right int) {
	n1 := (mid - left + 1)
	n2 := (right - mid)

	L := make([]int, n1)
	R := make([]int, n2)

	for i := 0; i < n1; i++ {
		L[i] = v[left+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = v[mid+1+j]
	}

	i, j, k := 0, 0, left

	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			v[k] = L[i]
			i++
		} else {
			v[k] = R[j]
			j++
		}
		k++
	}

	for i < n1 {
		v[k] = L[i]
		i++
		k++
	}

	for j < n2 {
		v[k] = R[j]
		j++
		k++
	}
}

func mergeSort(v []int, left, right int) {
	if left >= right {
		return
	}

	mid := left + (right-left)/2
	mergeSort(v, left, mid)
	mergeSort(v, mid+1, right)
	merge(v, left, mid, right)
}

func printSolution(v []int) {
	for _, value := range v {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Number of elements: ")
	fmt.Scan(&n)

	v := make([]int, n)
	fmt.Print("Elements: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	mergeSort(v, 0, n-1)

	fmt.Println("Sorted vector:")
	printSolution(v)
}
