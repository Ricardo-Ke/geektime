package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Go 爱好者"
	fmt.Printf("%q\n", str)
	fmt.Printf("runes(char): %q\n", []rune(str))
	fmt.Printf("runes(hex): %x\n", []rune(str))
	fmt.Printf("bytes(hex): % x\n", []byte(str))

	for i, c := range str {
		fmt.Printf("%d: %q [%b]\n", i, c, []byte(string(c)))
	}

	for i := 0; i < utf8.RuneCountInString(str); i++ {
		// ch, width := utf8.DecodeRuneInString(str)

		println(i)
	}
}
