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
	_
	_
	_
	// data type
	ID
	NUM
	_
	// op type
	ADD    // +
	SUB    // -
	MUL    // *
	QUO    // /
	COLON  // :
	ASSIGN // :=
	_
	LSS // <
	NEQ // <>
	LEQ // <=
	GTR // >
	GEQ // >=
	EQ  // =

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

// NewInput returns the object for program's position record
func NewInput(program string) *Input {
	if program[len(program)-1] != '#' {
		program += "#"
	}
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

// SkipWhitespace will skip ' \t\n'
func (i *Input) SkipWhitespace() {
	for i.program[i.position] == ' ' || i.program[i.position] == '\t' || i.program[i.position] == '\n' {
		i.position++
	}
}

// Next returns the next char
func (i *Input) Next() (byte, error) {
	if i.EOF() {
		return '#', errors.New("EOF")
	}
	i.position++
	if i.program[i.position] == '\n' {
		i.row++
		i.col = 0
	} else {
		i.col++
	}
	return i.program[i.position], nil
}

// Collapse returns error message for lex error
func (i *Input) Collapse(msg string) error {
	return fmt.Errorf("Error:%s row:%d col:%d", msg, i.row, i.col)
}

// IsLetter returns true if ch is letter
func IsLetter(ch byte) bool {
	if ch > 'a' && ch < 'z' || ch > 'A' && ch < 'Z' || ch == '_' || ch == '$' {
		return true
	}
	return false
}

// IsDigit returns true if ch is a digit
func IsDigit(ch byte) bool {
	if ch > '0' && ch < '9' {
		return true
	}
	return false
}

// func main() {
// 	program := "begin x:=9; if x>9 then x:=2*x+1/3; end #"
// 	input := NewInput(program)
// 	ch2 := input.Peek()
// 	ch1, _ := input.Next()
// 	fmt.Println(ch1 == ch2)

// 	if ch := 1; ch == 2 {

// 	} else if ch2 := 3; ch2 == 3 {
// 		fmt.Println(ch)
// 	}

// 	for k, v := range tokens {
// 		fmt.Println(k, v)
// 	}
// }
