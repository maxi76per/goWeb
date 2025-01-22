package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	seen := make(map[rune]bool)
	for _, char := range strings.ToLower(s) {
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}

func main() {
	fmt.Println(isUnique("abcd"))      // Output: true
	fmt.Println(isUnique("abCdefAaf")) // Output: false
	fmt.Println(isUnique("aabcd"))     // Output: false
}