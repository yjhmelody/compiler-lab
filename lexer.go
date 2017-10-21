package main

import (
	"errors"
	"fmt"
	"strconv"
)

// keywords: begin  if  then  while  do  end
// operator: :  :=  +  -  *  /  <  <=  <>  >  >=  =  ;  (  )  #
// ID = letter (letter | digit)*
// NUM = digit digit*
// 空格有空白、制表符和换行符组成。空格一般用来分隔ID、NUM、运算符、界符和关键字，词法分析阶段通常被忽略

// Token is the set of lexical tokens
type Token int

// String returns the string corresponding to the token tok.
func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

// begin x:=9; if x>9 then x:=2*x+1/3; end #
const (
	// SHARP is '#' which is the end char
	SHARP Token = iota
	// keywords
	BEGIN
	IF
	THEN
	WHILE
	DO
	END
	// data type
	ID
	NUM
	// op type
	ADD    // +
	SUB    // -
	MUL    // *
	QUO    // /
	COLON  // :
	ASSIGN // :=
	LSS    // <
	NEQ    // <>
	LEQ    // <=
	GTR    // >
	GEQ    // >=
	EQ     // =

	SEMCOLON // ;
	LPAREN   // (
	RPAREN   // )
)

var tokens = [...]string{
	SHARP:    "#",
	BEGIN:    "begin",
	IF:       "if",
	THEN:     "then",
	WHILE:    "while",
	DO:       "do",
	END:      "end",
	ID:       "id",
	NUM:      "num",
	ADD:      "+",
	SUB:      "-",
	MUL:      "*",
	QUO:      "/",
	COLON:    ":",
	ASSIGN:   ":=",
	LSS:      "<",
	NEQ:      "<>",
	LEQ:      "<=",
	GTR:      ">",
	GEQ:      ">=",
	EQ:       "=",
	SEMCOLON: ";",
	LPAREN:   "(",
	RPAREN:   ")",
}

func Lex() {

}

var keywords map[string]Token

// init will be called before main function
func init() {
	keywords = make(map[string]Token)
	for i := BEGIN; i <= END; i++ {
		keywords[tokens[i]] = i
	}
}

// Input records the lex position
type Input struct {
	position, row, col int
	program            string
}

// NewInput returns the object for program's record
func NewInput(program string) *Input {
	return &Input{
		position: 0,
		row:      1,
		col:      1,
		program:  program,
	}
}

// EOF returns true when program gets to the end
func (i *Input) EOF() bool {
	return i.position >= len(i.program)
}

// Peek returns current position char
func (i *Input) Peek() byte {
	if i.EOF() {
		return '#'
	}
	return i.program[i.position]
}

func (i *Input) skipWhitespace() {

}

// Next returns the next char
func (i *Input) Next() (byte, error) {
	if i.EOF() {
		return '#', errors.New("EOF")
	}
	if i.program[i.position] == '\n' {
		i.row++
		i.col = 0
	} else {
		i.col++
	}
	char := i.program[i.position]
	i.position++
	return char, nil
}

// Collapse returns error message for lex error
func (i *Input) Collapse(msg string) error {
	return fmt.Errorf("Error:%s row:%d col:%d", msg, i.row, i.col)
}

func main() {
	// program := "begin x:=9; if x>9 then x:=2*x+1/3; end #"

}