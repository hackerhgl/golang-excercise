package main

import "fmt"

func main() {
	shuffle := []int{1, 1, 4, 1, 99, 4, 3, 8, 9, 3}

	size := len(shuffle)

	start := 0

	fmt.Println("shuffled array", shuffle)

	for true {
		nextIndex := start + 1
		if nextIndex >= size {
			break
		}
		next := shuffle[nextIndex]
		curr := shuffle[start]

		if curr > next {
			shuffle[start] = next
			shuffle[nextIndex] = curr
			start = 0
		} else {
			start++
		}

	}

	fmt.Println("sorted array", shuffle)
}
