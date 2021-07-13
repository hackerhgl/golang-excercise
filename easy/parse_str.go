package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "11:01:00PM"
	size := len(s)
	d := strings.Split(s[:size-2], ":")
	day := s[size-2 : size]

	if day == "AM" && d[0] == "12" {

		d[0] = "00"
	} else if day == "PM" && d[0] != "12" {
		tmp, _ := strconv.Atoi(d[0])
		d[0] = strconv.Itoa(tmp + 12)
	}

	fmt.Println(strings.Join(d, ":"))

}
