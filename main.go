package main

import (
	"fmt"
	"new-lang/lexer"
	"os"
	"log"
)

func main() {
	// debug print
	fmt.Printf("hello")

	// panファイルを読み込みたい
	file, err := os.Open("pan.pan")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// lexerを呼び出す
	lexer.Lexer()
}
