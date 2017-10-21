package main

import (
	"log"
)

// Scanner stores token
type Scanner struct {
	input Input
	token string
	syn   Token
}

// NewScanner returns Scanner pointer
func NewScanner(input Input) *Scanner {
	return &Scanner{
		input: input,
		token: "",
	}
}

// EOF returns true if token is null
func (s *Scanner) EOF() bool {
	if s.token == "" {
		return true
	}
	return false
}

// Peek returns current <token, syn>
func (s *Scanner) Peek() string {
	if s.token != "" {
		return s.token
	}
	return s.Next()
}

// Next returns next <token, syn>
func (s *Scanner) Next() string {

}

func (s *Scanner) setLex(token string, syn Token) {
	s.token = token
	s.syn = syn
}

func (s *Scanner) read() {
	s.input.SkipWhitespace()
	if s.input.EOF() {
		s.token = "#"
	}
	ch := s.input.Peek()
	if IsDigit(ch) && ch != '0' {
		s.readNum()
		return
	} else if IsLetter(ch) {
		s.readID()
		return
	}
	switch ch {
	case '+', '-', '*', '/', ';', '(', ')':
		s.token = string(ch)
		return
	default:
		s.readOp()
	}

}

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

func (s *Scanner) readOp() {
	// := < <> <= > >= = ; ( )
	if ch, _ := s.input.Next(); ch == ':' {
		if s.input.Peek() == '=' {
			s.setLex(":=", ASSIGN)
		}
	} else if ch == '=' {
		s.setLex("=", EQ)
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
	} else {
		log.Fatal("readOp error")
	}
}
