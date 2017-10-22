package main

import (
	"fmt"
	"log"
)

// Scanner stores token
type Scanner struct {
	input *Input
	token string
	syn   Token
}

// NewScanner creates a scanner to scan token
func NewScanner(input *Input) *Scanner {
	return &Scanner{
		input: input,
		token: "",
	}
}

// SkipWhitespace will skip ' \t\n'
func (s *Scanner) SkipWhitespace() {
	ch, _ := s.input.Next()
	for ch == ' ' || ch == '\t' || ch == '\n' {
		ch, _ = s.input.Next()
	}
}

// EOF returns true if token is '#'
func (s *Scanner) EOF() bool {
	if s.token == "#" {
		return true
	}
	return false
}

// Peek returns current <token, syn>
func (s *Scanner) Peek() (string, Token) {
	return s.token, s.syn
}

// Next returns the next <token, syn>
func (s *Scanner) Next() (string, Token) {
	s.read()
	return s.Peek()
}

func (s *Scanner) setLex(token string, syn Token) {
	s.token = token
	s.syn = syn
}

// read chars until gets a total token
func (s *Scanner) read() {
	// s.input.SkipWhitespace()
	ch := s.input.Next()
	if ch := s.input.Next(); ch == '#' {
		s.setLex("#", SHARP)
	} else if IsDigit(ch) {
		s.readNum()
	} else if IsLetter(ch) {
		s.readID()
	} else {
		s.readOp()
	}
	s.SkipWhitespace()

}

// read the num type
func (s *Scanner) readNum() {
	// first digit [1-9]
	ch, _ := s.input.Next()
	s.token = string(ch)
	// digit [0-9]
	for {
		ch, _ := s.input.Next()
		if IsDigit(ch) {
			s.token += string(ch)
		} else {
			break
		}
	}
	s.syn = NUM
}

// read the identifier and keywords
func (s *Scanner) readID() {
	// first letter [A-Za-z_$]
	ch, _ := s.input.Next()
	s.token = string(ch)
	// letter [A-Za-z0-9_$]
	for {
		ch, _ = s.input.Next()
		if IsLetter(ch) || IsDigit(ch) {
			s.token += string(ch)
		} else {
			break
		}
	}
	// recognize keywords
	if v, ok := keywords[s.token]; ok {
		s.syn = v
	} else {
		s.syn = ID
	}
}

// read the operations
func (s *Scanner) readOp() {
	// operator: :  :=  +  -  *  /  <  <=  <>  >  >=  =  ;  (  )  #
	if ch, _ := s.input.Next(); ch == ':' {
		if s.input.Peek() == '=' {
			s.setLex(":=", ASSIGN)
		}
	} else if ch == '=' {
		s.setLex("=", EQ)
	} else if ch == '+' {
		s.setLex("+", ADD)
	} else if ch == '-' {
		s.setLex("-", SUB)
	} else if ch == '*' {
		s.setLex("*", MUL)
	} else if ch == '/' {
		s.setLex("/", QUO)
	} else if ch == ';' {
		s.setLex(";", SEMCOLON)
	} else if ch == '<' {
		switch s.input.Peek() {
		case '>':
			s.setLex("<>", NEQ)
		case '=':
			s.setLex("<=", LEQ)
		default:
			s.setLex("<", LSS)
		}
	} else if ch == '>' {
		switch s.input.Peek() {
		case '=':
			s.setLex(">=", GEQ)
		default:
			s.setLex(">", GTR)
		}
	} else if ch == ';' {
		s.setLex(";", SEMCOLON)
	} else if ch == '(' {
		s.setLex("(", LPAREN)
	} else if ch == ')' {
		s.setLex(")", RPAREN)
	} else if ch == '#' {
		s.setLex("#", SHARP)
	} else {
		log.Fatal("readOp error")
	}
}

func main() {
	program := "begin x := 9; if x>9 then x:=2*x+1/3; end #"
	// program := " begin x 09 then x ; end #"
	scanner := NewScanner(NewInput(program))
	token, syn := scanner.Next()
	for !scanner.EOF() {
		token, syn = scanner.Next()
		fmt.Printf("<%s,%d>\n", token, syn)
		fmt.Println(scanner.input.position)
	}
	fmt.Printf("<%s,%d>\n", token, syn)
}
