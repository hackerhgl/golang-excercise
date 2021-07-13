package main

import "fmt"

func main() {
	candles := []int32{3, 2, 1, 3, 3, 3}

	max := int32(0)
	size := len(candles)
	counts := make(map[int32]int32)
	index := 0
	for true {
		if index >= size {
			break
		}
		curr := int32(candles[index])
		if curr > max {
			max = curr
		}
		index++
		if counts[curr] == 0 {
			counts[curr] = 1
		} else {
			counts[curr]++
		}
	}

	fmt.Println(counts[max])
}
