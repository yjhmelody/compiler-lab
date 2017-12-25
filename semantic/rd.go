package main

import (
	"fmt"
	"strconv"

	"github.com/yjhmelody/compiler-lab/lexer"
	"github.com/yjhmelody/compiler-lab/stack"
)

var scanner *lexer.Scanner
var table = &stack.Stack{}
var offsets = &stack.Stack{}
var symbolTable *SymbolTable

// var globalSymbols map[lexer.Token]map[string]string
var Tvalue map[string]string

func init() {
	fmt.Println("init")
	Tvalue = make(map[string]string)
	table = stack.NewStack()
	offsets = stack.NewStack()
}

// Parse
func Parse(s *lexer.Scanner) bool {
	scanner = s
	return parse_P()
}

func parse_P() bool {
	parse_M()
	parse_D()
	st, ok := table.Peak().(*SymbolTable)
	if !ok {
		fmt.Println("stack error", table.Len())
		return false
	}
	offset, ok := offsets.Peak().(int)
	if !ok {
		fmt.Println("stack error", offsets.Len())
		return false
	}

	// addwidth(top(tblptr),top(offset));
	// pop(tblptr);  pop(offset)
	st.addWidth(offset)
	table.Pop()
	offsets.Pop()
	return true
}

func parse_M() {
	// t = mktable(nil); push(t,tblptr); push(0,offset)
	symbolTable = mktable("", nil)
	table.Push(symbolTable)
	// 0
	offsets.Push(0)
}

func parse_D() bool {

	// id : T D2
	if tok, syn := scanner.Next(); syn == lexer.ID {
		if _, syn := scanner.Next(); syn == lexer.COLON {

			parse_T()

			st, ok := table.Peak().(*SymbolTable)
			if !ok {
				fmt.Println("stack error", table.Len())
				return false
			}
			offset, ok := offsets.Peak().(int)
			if !ok {
				fmt.Println("stack error", offsets.Len())
				return false
			}
			// Tvalue, ok := globalSymbols[T].(map[string]string)
			// Tvalue := globalSymbols[T]
			// Tvalue := TTable
			if !ok {
				fmt.Println("globalSymbols[T] error")
				return false
			}

			// enter(top(tblptr),id.name,T.type,top(offset));
			// top(offset) = top(offset) + T.width
			st.enter(tok, Tvalue["type"], offset)
			offsets.Pop()
			width, err := strconv.Atoi(Tvalue["width"])
			if err != nil {
				fmt.Println("err")
			}
			offsets.Push(width + offset)
			parse_D2()
		}
	}
	return false
}

func parse_T() string {
	tok, _ := scanner.Next()

	// T.type = integer; T.width = 4
	// T.type = real; T.width = 8
	// T.type = pointer(T1.type); T.width = 4
	switch tok {
	case "integer":
		Tvalue["type"] = tok
		Tvalue["width"] = "4"
	case "real":
		Tvalue["type"] = tok
		Tvalue["width"] = "8"
	case "ptr":
		Tvalue["type"] = tok + " " + parse_T()
		// Tvalue["type"] = tok + T2["type"]
		Tvalue["width"] = "4"
	default:
		fmt.Println("T error")
		panic("parse_T")
	}

	return tok
}

func parse_D2() {
	if _, syn := scanner.Next(); syn == lexer.SEMCOLON {
		parse_D()
		parse_D2()
	}
}

func main() {
	program := `id1:real; id2:ptr integer; id3:integer; id4: real;`
	Parse(lexer.NewScanner(lexer.NewInput(program)))
	fmt.Println("width", symbolTable.width)
	for k, v := range symbolTable.symbols {
		fmt.Println(k, v)
	}
}
