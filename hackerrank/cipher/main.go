package main

import (
	"fmt"
	"strings"
)

func main() {

	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	//	fmt.Printf("length: %d\n", length)
	//	fmt.Printf("input: %s\n", input)
	//	fmt.Printf("delta: %d\n", delta)
	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ret := ""
	for _, ch := range input {
		switch {
		case strings.IndexRune(string(alphabetLower), ch) >= 0:
			ret = ret + string(rotate(ch, delta, []rune(alphabetLower)))
		case strings.IndexRune(string(alphabetUpper), ch) >= 0:
			ret = ret + string(rotate(ch, delta, []rune(alphabetUpper)))
		default:
			ret = ret + string(ch)
		}
	}
	fmt.Println(ret)
	//_ = alphabet

	//newRune := rotate('z', 2, alphabet)
	//fmt.Println(string(newRune))

}

func rotate(s rune, delta int, key []rune) rune {
	idx := strings.IndexRune(string(key), s)
	if idx < 0 {
		panic("idx < 0")
	}
	idx = (idx + delta) % len(key)

	return key[idx]
}
