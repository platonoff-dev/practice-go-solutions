package anagram

import "strings"

func isAnagram(word1, word2 string) bool {
	// Prepare strings for comparison
	word1 = strings.ToLower(strings.Replace(word1, " ", "", -1))
	word2 = strings.ToLower(strings.Replace(word2, " ", "", -1))

	// Check basic cases when strings is not anagrams
	if len(word1) != len(word2) || word1 == word2 {
		return false
	}

	// Count letters in both string. Strings are anagrams if final count for each letter is 0
	letterCounts := map[rune]int{}
	for _, ch := range word1 {
		letterCounts[ch] += 1
	}

	for _, ch := range word2 {
		if _, ok := letterCounts[ch]; !ok {
			return false
		} else {
			letterCounts[ch] -= 1
		}
	}

	for _, count := range letterCounts {
		if count != 0 {
			return false
		}
	}

	return true
}

func FindAnagrams(dictionary []string, word string) (result []string) {
	for _, dictWord := range dictionary {
		if isAnagram(word, dictWord) {
			result = append(result, dictWord)
		}
	}

	return
}
