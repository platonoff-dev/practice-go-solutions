package buildword

func SearchParts(word string, fragments []string) int {
	minLength := 0
	length := 0
	for _, fragment := range fragments {
		if len(fragment) > len(word) || word[:len(fragment)] != fragment {
			continue
		}

		if fragment == word {
			return 1
		}

		restLength := SearchParts(word[len(fragment):], fragments)
		if restLength == 0 {
			continue
		}

		length = restLength + 1
		if length != 0 && (length < minLength || minLength == 0) {
			minLength = length
		}
	}

	return minLength
}

func BuildWord(word string, fragments []string) int {
	return SearchParts(word, fragments)
}
