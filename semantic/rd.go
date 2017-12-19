package main

import (
	"fmt"
	"strconv"

	"github.com/yjhmelody/compiler-lab/lexer"
	"github.com/yjhmelody/compiler-lab/stack"
	// "github.com/yjhmelody/compiler-lab/stack"
)

var scanner *lexer.Scanner
var table = &stack.Stack{}
var offsets = &stack.Stack{}
var globalSymbols map[lexer.Token]interface{}

func init() {
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
	t := mktable("", nil)
	table.Push(t)
	offsets.Push(0)
}

func parse_D() bool {
	// if scanner.EOF() {
	// 	return false
	// }

	// id : T D2
	if _, syn := scanner.Next(); syn == lexer.ID {
		if tok, syn := scanner.Next(); syn == lexer.COLON {
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
			Tvalue, ok := globalSymbols[T].(map[string]string)
			if !ok {
				fmt.Println("globalSymbols[T] error")
				return false
			}

			// enter(top(tblptr),id.name,T.type,top(offset));
			// top(offset) = top(offset) + T.width
			st.enter(tok, Tvalue["type"], offset)
			offsets.Pop()
			wdith, err := strconv.Atoi(Tvalue["width"])
			if err != nil {
				fmt.Println("err")
			}
			offsets.Push(wdith + offset)

			parse_D2()
		}
	}
	return false
}

func parse_T() string {
	tok, _ := scanner.Next()
	Tvalue, ok := globalSymbols[T].(map[string]string)
	if !ok {
		fmt.Println("globalSymbols[T] error")
		panic("paser_T")
	}

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
		Tvalue["type"] = tok + parse_T()
		// Tvalue["type"] = tok + T2["type"]
		Tvalue["width"] = "4"
	default:
		fmt.Println("T error")
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
	program := `id1 : real ; id2:ptr integer; id3:integer`
	Parse(lexer.NewScanner(lexer.NewInput(program)))
}
