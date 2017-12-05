package main

import (
	"fmt"

	"github.com/yjhmelody/compiler-lab/LL1"
	"github.com/yjhmelody/compiler-lab/lexer"
)

func main() {
	program := `id +++ id3`
	// fmt.Scanf("%s", program)
	ok := LL1.Analysis(lexer.NewScanner(lexer.NewInput(program)))
	fmt.Println("recognize?", ok)
}
