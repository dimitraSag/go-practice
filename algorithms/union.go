package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]
	seen := ""

	contains := func(ch rune, str string) bool {
		for _, c := range str {
			if c == ch {
				return true
			}
		}
		return false
	}

	printUnique := func(s string) {
		for _, ch := range s {
			if !contains(ch, seen) {
				seen += string(ch)
				fmt.Printf("%c", ch)
			}
		}
	}

	printUnique(s1)
	printUnique(s2)
	fmt.Println()
}
