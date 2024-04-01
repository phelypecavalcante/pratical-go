package main

import (
	"fmt"
	"unicode/utf8"
)

func banner(text string, width int) {
	// BUG: len is in bytes
	// padding := (width - len(text)) / 2
	padding := (width - utf8.RuneCountInString(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}

	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}

	fmt.Println()

}

// isPalindrome("g") -> true
// isPalindrome("go") -> false
// isPalindrome("gog") -> true
// isPalindrome("gogo") -> false
func isPalindrome(s string) bool {
	rs := []rune(s)
	lastPos := utf8.RuneCountInString(s) - 1
	for i := 0; i < len(rs)/2; i++ {
		if rs[i] != rs[lastPos-i] {
			return false
		}
	}
	return true
}

func main() {
	banner("Go", 6)
	banner("G☺", 6)

	s := "G☺"

	fmt.Println("len:", len(s))
	// code point = run ~= unicode character
	for i, r := range s {
		fmt.Printf("%d - %c of type %T\n", i, r, r)
		// rune (int32)
	}

	// byte (uint8)
	// rune (int32)

	b := s[0]
	fmt.Printf("%c of type %T\n", b, b)

	b = s[1]
	fmt.Printf("%c of type %T\n", b, b)

	x, y := 1, "1"
	fmt.Printf("x=%#v, y=%#v\n", x, y) // Use #v in debug/log

	fmt.Printf("%20s!\n", s)

	fmt.Println("g", isPalindrome("g"))
	fmt.Println("go", isPalindrome("go"))
	fmt.Println("gog", isPalindrome("gog"))
	fmt.Println("gogo", isPalindrome("gogo"))
	fmt.Println("g☺g", isPalindrome("g☺g"))
}
