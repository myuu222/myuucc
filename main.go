package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "引数の個数が正しくありません")
		os.Exit(1)
	}

	var path = os.Args[1]
	tokenizer = NewTokenizer()
	tokenizer.Tokenize(path)
	Parse()
	pipeline(code)

	// アセンブリの前半部分
	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".globl main")

	fmt.Println(".data")
	for _, str := range Env.StringLiterals {
		fmt.Println(str.label + ":")
		fmt.Println("  .string " + str.value)
	}
	fmt.Println(".text")

	for _, c := range code {
		// 抽象構文木を下りながらコード生成
		gen(c)
	}
}
