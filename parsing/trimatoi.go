package piscine

func TrimAtoi(s string) int {
	sign := 1
	result := 0
	signSet := false
	numberStarted := false
	for _, char := range s {
		if char == '-' && !signSet && !numberStarted {
			sign = -1
			signSet = true
		} else if char == '+' && !signSet && !numberStarted {
			signSet = true
		} else if char >= '0' && char <= '9' {
			toint := int(char - '0')
			result = result*10 + toint
			numberStarted = true
		}
	}
	return sign * result
}
