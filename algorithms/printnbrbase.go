package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		z01.PrintRune('\n')
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		nbr *= -1
	}

	baseLen := len(base)
	var result string
	for nbr > 0 {
		remainder := nbr % baseLen
		result = string(base[remainder]) + result
		nbr /= baseLen
	}

	if result == "" {
		z01.PrintRune('0')
	} else {
		for _, r := range result {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	chars := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' {
			return false
		}
		if chars[r] {
			return false
		}
		chars[r] = true
	}
	return true
}
