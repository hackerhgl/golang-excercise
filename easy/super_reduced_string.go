package main

import "fmt"

func main() {
	str := "aa"

	// fmt.Println(superReducedString("aaabccddd"))
	start := 0
	for true {
		size := len(str)
		nextIndex := start + 1

		if nextIndex >= size {
			break
		}

		next := str[nextIndex]
		curr := str[start]

		if curr == next {
			r := []rune(str)
			r = delChar(r, nextIndex)
			r = delChar(r, start)
			str = string(r)
			start = 0
			continue
		}
		start++

	}
	if str == "" {
		fmt.Println("Empty String")
	} else {

		fmt.Println(str)
	}

}

func superReducedString(s string) string {
	start := 0

	for true {
		size := len(s)
		nextIndex := start + 1

		if nextIndex >= size {
			break
		}

		next := s[nextIndex]
		curr := s[start]

		if curr == next {
			r := []rune(s)
			r = delChar(r, nextIndex)
			r = delChar(r, start)
			s = string(r)
			start = 0
			continue
		}
	}
	return s
}

func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}
