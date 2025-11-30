package lexer

import "fmt"

func Lexer(input string) {
	fmt.Printf("lexer input: %s\n", input)

	for i, ch := range input {
		fmt.Printf("position: %d, ch: %c\n", i, ch)
	}
}

