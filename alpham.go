package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	input := os.Args[1]
	fmt.Println(alphamirror(input))
}

func alphamirror(s string) string {
	result := []byte{}
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'a' && ch <= 'z' {

			opposite := 'z' - (ch - 'a')
			result = append(result, opposite)
		} else if ch >= 'A' && ch <= 'Z' {

			opposite := 'Z' - (ch - 'A')
			result = append(result, opposite)
		} else {

			result = append(result, ch)
		}
	}
	return string(result)
}
