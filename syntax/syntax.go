package main

import (
	"github.com/yjhmelody/compiler-lab/lexer"
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
