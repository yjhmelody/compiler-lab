package main

import (
	"fmt"

	"github.com/yjhmelody/compiler-lab/lexer"
	"github.com/yjhmelody/compiler-lab/stack"
)

// Right stores the production's right part
type Right []lexer.Token

// Items store the tokens corresponds to the production
type Items map[lexer.Token]Right

// AnalysisTable stores the table for predicting LL(1) grammar
type AnalysisTable map[lexer.Token]Items

// defined for syntax
const (
	EPISILON lexer.Token = lexer.EPISILON + iota
	E
	E2
	T
	T2
	F
)

var analysisTable = AnalysisTable{
	E: {
		lexer.ID:     Right{T, E2},
		lexer.ADD:    nil,
		lexer.MUL:    nil,
		lexer.LPAREN: Right{T, E2},
		lexer.RPAREN: nil,
		lexer.SHARP:  nil,
	},
	E2: {
		lexer.ID:     nil,
		lexer.ADD:    Right{lexer.ADD, T, E2},
		lexer.MUL:    nil,
		lexer.LPAREN: nil,
		lexer.RPAREN: Right{EPISILON},
		lexer.SHARP:  Right{EPISILON},
	},
	T: {
		lexer.ID:     Right{F, T2},
		lexer.ADD:    nil,
		lexer.MUL:    nil,
		lexer.LPAREN: Right{F, T2},
		lexer.RPAREN: nil,
		lexer.SHARP:  nil,
	},
	T2: {lexer.ID: nil,
		lexer.ADD:    Right{EPISILON},
		lexer.MUL:    Right{lexer.ADD, F, T2},
		lexer.LPAREN: nil,
		lexer.RPAREN: Right{EPISILON},
		lexer.SHARP:  Right{EPISILON},
	},
	F: {
		lexer.ID:     Right{lexer.ID},
		lexer.ADD:    nil,
		lexer.MUL:    nil,
		lexer.LPAREN: Right{lexer.LPAREN, E, lexer.RPAREN},
		lexer.RPAREN: nil,
		lexer.SHARP:  nil,
	},
}

// var program = `9x9x
// 0099
// ??$$
// ++
// begin 9x:=?$00999; if x%><<>9 t99he&n x:=2**x+1/3; end #
// `

// scanner := NewScanner(NewInput(program))
// token, syn := scanner.Next()
// fmt.Printf("<'%s', %d>\n", token, syn)
// for !scanner.EOF() {
// 	token, syn = scanner.Next()
// 	fmt.Printf("<'%s', %d>\n", token, syn)
// }

// Analysis uses the scanner to recognize LL(1) grammar
func Analysis(s *lexer.Scanner) {
	stack := stack.NewStack()
	stack.Push(lexer.SHARP)
	stack.Push(E)

	X, ok := stack.Peak().(lexer.Token)
	if !ok {
		fmt.Println("error")
	}
	// when stack is not empty
	for X != lexer.SHARP {
		_, syn := s.Next()
		if X == syn {
			stack.Pop()
		} else if X < lexer.EPISILON {
			// X is a terminal
			fmt.Println("a terminal")
		} else if v, ok := analysisTable[X][syn]; v == nil || !ok {
			fmt.Println("table error")
		} else if v, ok := analysisTable[X][syn]; ok {

		}
	}
}
