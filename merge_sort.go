package main

import (
	"fmt"
	"time"
)

func Merge(a []int, i int, j int, m int) {
	c := []int{}
	k := 0
	l := m + 1

	for i <= m && l <= j {
		if a[i] > a[l] {
			c = append(c, a[l])
			l++
		} else {
			c = append(c, a[i])
			i++
		}
		k++
	}

	if i <= m {
		for i <= m {
			c = append(c, a[i])
			i++
			k++
		}
	}
	if l <= j {
		for l <= j {
			c = append(c, a[l])
			l++
			k++
		}
	}
	for k > 0 {
		a[j] = c[k-1]
		k--
		j--
	}
}

func MergeSort(arr []int, i int, j int) {
	if i == j {
		return
	}
	mid := i + (j-i)/2
	MergeSort(arr, i, mid)
	MergeSort(arr, mid+1, j)
	Merge(arr, i, j, mid)
}

func main() {
	var array = []int{5, 3, 4, 5, 6, 2, 1, 3, 0, 2, 1, 77, 3, 2, 1, 23, 3, 21, 33, 2, 1, 1, 23, 2, 211, 222, 2,1000000000000000000,1000}
	start := time.Now()
	MergeSort(array, 0, len(array)-1)
	fmt.Printf("Time taken %f", time.Since(start).Seconds())
	fmt.Println(array)
}
