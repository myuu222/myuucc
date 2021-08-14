package main

import (
	"fmt"
	"os"

	"github.com/myuu222/myuugo/codegen"
	"github.com/myuu222/myuugo/parse"
	"github.com/myuu222/myuugo/passes"
)

func usage() {
	fmt.Fprintln(os.Stderr, "引数の個数が正しくありません")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	var path = os.Args[1]
	var program = parse.Parse(path)

	passes.Semantic(program)

	codegen.GenX86_64(program)
}
