package piscine

func myTrimSpace(s string) string {
	start := 0
	end := len(s) - 1

	for start <= end && (s[start] == ' ' || s[start] == '\t' || s[start] == '\n' || s[start] == '\r') {
		start++
	}

	for end >= start && (s[end] == ' ' || s[end] == '\t' || s[end] == '\n' || s[end] == '\r') {
		end--
	}

	if start > end {
		return ""
	}

	return s[start : end+1]
}

func myFields(s string) []string {
	words := []string{}
	wordStart := -1

	for i, ch := range s {
		if ch != ' ' && ch != '\t' && ch != '\n' && ch != '\r' {
			
			if wordStart == -1 {
				wordStart = i
			}
		} else {
		
			if wordStart != -1 {
				words = append(words, s[wordStart:i])
				wordStart = -1
			}
		}
	}

	
	if wordStart != -1 {
		words = append(words, s[wordStart:])
	}

	return words
}

func WordFlip(str string) string {
	trimmed := myTrimSpace(str)

	if trimmed == "" {
		return "Invalid Output\n"
	}

	words := myFields(trimmed)


	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	
	result := ""
	for i, w := range words {
		if i > 0 {
			result += " "
		}
		result += w
	}

	return result + "\n"
}
