package lexer

import (
	"fmt"
	"unicode"
)

type Token struct {
	Value rune
	Type string
}

// 次のトークンに進む
// 返し値: 次のトークンがあればtrue、なければfalse
func Tokenize(content string) []Token {
	fmt.Printf("call advance\n")
	fmt.Printf("content: %s\n", content)

	var tokens []Token
	for _, v := range content {
		// tokenを生成する
		t := token(v)
		// tokensにtokenをappendする
		tokens = append(tokens, t)
	}

	fmt.Printf("tokens: %+v\n", tokens)
	return tokens
}

func token(r rune) Token {
	if unicode.IsDigit(r) {
		return Token{Value: r, Type: "digit"}
	}
	return Token{Value: r, Type: "other"}
}
