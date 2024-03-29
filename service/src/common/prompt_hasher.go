package common

import (
	"crypto/sha512"
	"encoding/hex"
	"sort"
	"strings"
)

func GetPromptHash(str string) (string, string) {
	// Step 1: Replace consecutive spaces with a single space
	str = strings.Join(strings.Fields(str), " ")

	// Step 2: Split the string into an array of words
	arr := strings.Split(str, ",")

	// Step 3: Change all capitalized words to lowercase
	for i, word := range arr {
		arr[i] = strings.ToLower(strings.TrimSpace(word))
	}

	// Step 4: Make everything with multiple spaces into one and remove duplicates
	uniqueWords := make(map[string]bool)
	for _, word := range arr {
		if len(word) == 0 {
			continue
		}
		uniqueWords[word] = true
	}

	// Step 5: Convert the map of unique words to a slice of words and sort it
	sortedWords := make([]string, 0, len(uniqueWords))
	for word := range uniqueWords {
		sortedWords = append(sortedWords, word)
	}
	sort.Strings(sortedWords)

	// Step 6: Concatenate the words into a single string
	words := strings.Join(sortedWords, ", ")

	// Step 7: Calculate the SHA512 hash of the concatenated string
	hashBytes := sha512.Sum512([]byte(words))
	hash := hex.EncodeToString(hashBytes[:])

	// Step 8: Return the concatenated string and the hash
	return words, hash
}
