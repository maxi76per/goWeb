package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	str := "главрыба"
	reversed := reverseString(str)
	fmt.Println(reversed) // Output: абырвалг
}