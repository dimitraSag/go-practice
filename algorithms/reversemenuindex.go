package piscine

func ReverseMenuIndex(menu []string) []string {
	list2 := make([]string, len(menu))
	for i := 0; i < len(menu); i++ {
		list2[i] = menu[len(menu)-1-i]
	}
	return list2
}
