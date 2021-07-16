package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag := true

	reader := bufio.NewReader(os.Stdin)

	for flag {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		notEmpty := len(text) > 1

		isQuestion := notEmpty && strings.Contains(text, "?")
		isYelling := notEmpty && text == strings.ToUpper(text)

		if isQuestion && isYelling {
			fmt.Println("Calm down, I know what I'm doing!")
		} else if isQuestion {
			fmt.Println("Sure.")
		} else if isYelling {
			fmt.Println("Whoa, chill out!")
		} else if !notEmpty {
			fmt.Println("Fine. Be that way!")
		} else {
			fmt.Println("Whatever.")
		}

	}
}
