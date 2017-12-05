package LL1

import (
	"fmt"

	"github.com/yjhmelody/compiler-lab/lexer"
	"github.com/yjhmelody/compiler-lab/stack"
)

// Syntax
// E -> T E2
// E2 -> + T E2 | ε
// T -> F T2
// T2 -> * F T2 | ε
// F -> ( E ) | i

// Right stores the production's right part
type Right []lexer.Token

// Items store the tokens corresponds to the production
type Items map[lexer.Token]Right

// AnalysisTable stores the table for predicting LL(1) grammar
type AnalysisTable map[lexer.Token]Items

// defined for syntax
const (
	// 29
	EPISILON lexer.Token = lexer.EPISILON + iota
	E
	E2
	T
	T2
	F
)

// when meets nil, analysis will warning and ignore some syntax error
var analysisTable = AnalysisTable{
	E: {
		lexer.ID: Right{T, E2},
		// lexer.ADD:    nil,
		// lexer.MUL:    nil,
		lexer.LPAREN: Right{T, E2},
		lexer.RPAREN: nil,
		lexer.SHARP:  nil,
	},
	E2: {
		lexer.ADD:    Right{lexer.ADD, T, E2},
		lexer.RPAREN: Right{EPISILON},
		lexer.SHARP:  Right{EPISILON},
	},
	T: {
		lexer.ID:     Right{F, T2},
		lexer.ADD:    nil,
		lexer.LPAREN: Right{F, T2},
		lexer.RPAREN: nil,
		lexer.SHARP:  nil,
	},
	T2: {
		lexer.ADD:    Right{EPISILON},
		lexer.MUL:    Right{lexer.MUL, F, T2},
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

// Analysis uses the scanner to recognize LL(1) grammar
func Analysis(s *lexer.Scanner) bool {
	flag := true
	stack := stack.NewStack()
	stack.Push(lexer.SHARP)
	stack.Push(E)

	X, ok := stack.Peak().(lexer.Token)
	if !ok {
		fmt.Println("stack error")
	}
	_, curTok := s.Next()
	// when stack is not empty
	for stack.Peak() != lexer.SHARP {
		if X == curTok {
			fmt.Println("matched:", X)
			_, curTok = s.Next()
			stack.Pop()
		} else if X < lexer.EPISILON {
			// when X is a terminal
			fmt.Println("terminal error:", X, curTok)
			flag = false
		} else if _, ok := analysisTable[X][curTok]; !ok {
			fmt.Println("ignore the token:", X, curTok)
			_, curTok = s.Next()
			flag = false
		} else if v := analysisTable[X][curTok]; v == nil {
			fmt.Println("synch error:", X, curTok)
			stack.Pop()
			flag = false
		} else if v, ok := analysisTable[X][curTok]; ok {
			fmt.Println("production:", X, "->", v)
			stack.Pop()
			// push the production to stack
			for i := len(v) - 1; i >= 0; i-- {
				if v[i] != EPISILON {
					stack.Push(v[i])
				}
			}
		}
		X, ok = stack.Peak().(lexer.Token)
		if !ok {
			fmt.Println("stack error", stack.Len())
			return false
		}
	}
	return flag
}
