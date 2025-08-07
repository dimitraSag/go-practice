package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	length := max - min
	result := make([]int, length) // Create a slice with the required length
	for i := 0; i < length; i++ {
		result[i] = min + i
	}
	return result
}
