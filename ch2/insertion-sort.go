package main

import (
	"fmt"
)

func main() {
	a := []int{5, 2, 4, 6, 1, 3}

	for j := 1; j < len(a); j++ {
		key := a[j]
		// insert a[j] into the sorted sequence a[1..j-1]
		i := j - 1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i = i - 1
		}
		a[i+1] = key
		fmt.Println("%v", a)
	}
}
