package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "saveChangesInTheEditor"

	arr := strings.Split(s, "")
	count := 1

	for _, c := range arr {
		if strings.ToUpper(c) == c {
			count++
		}
	}
	fmt.Println(count)
}
