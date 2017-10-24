package main

import (
	"fmt"
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
	s.input.SkipWhitespace()
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

// setLex records current token info
func (s *Scanner) setLex(token string, syn Token) {
	s.token = token
	s.syn = syn
}

// read chars until gets a total token
func (s *Scanner) read() {
	s.SkipWhitespace()
	if ch := s.input.Peek(); ch == '#' {
		s.setLex("#", SHARP)
	} else if IsDigit(ch) {
		s.readNum()
	} else if IsLetter(ch) {
		s.readID()
	} else if IsOpChar(ch) {
		if err := s.readOp(); err != nil {
			fmt.Println(err)
		}
	} else {
		s.read()
	}
}

// readNum read the num type
func (s *Scanner) readNum() {
	// first digit [1-9]
	ch := s.input.Peek()
	if ch == '0' {
		s.setLex("0", NUM)
		s.input.Next()
	} else {
		// digit [0-9]
		str := string(ch)
		for {
			ch, _ := s.input.Next()
			if IsDigit(ch) {
				str += string(ch)
			} else {
				break
			}
		}
		s.setLex(str, NUM)
	}
}

// readID read the identifier and keywords
func (s *Scanner) readID() {
	// first letter [A-Za-z_$]
	ch := s.input.Peek()
	str := string(ch)

	// letter [A-Za-z0-9_$]
	for {
		ch, _ = s.input.Next()
		if IsLetter(ch) || IsDigit(ch) {
			str += string(ch)
		} else {
			break
		}
	}

	// recognize keywords
	if kw, ok := keywords[str]; ok {
		s.setLex(str, kw)
	} else {
		s.setLex(str, ID)
	}
}

// readOp read the operations
func (s *Scanner) readOp() error {
	// operator: :  :=  +  -  *  /  <  <=  <>  >  >=  =  ;  (  )  #
	var tempByte []byte

	switch ch := s.input.Peek(); ch {
	case '=', '+', '-', '*', '/', ';', '(', ')', '#':
		s.setLex(string(ch), keywords[string(ch)])
		s.input.Next()
	case ':', '<', '>':
		tempByte = append(tempByte, ch)
		// switch
		switch ch, _ := s.input.Next(); ch {
		case '=', '>':
			tempByte = append(tempByte, ch)
			tempChar := string(tempByte)
			if _, ok := keywords[tempChar]; ok {
				s.setLex(tempChar, keywords[tempChar])
				s.input.Next()
			}
		default:
			tempChar := string(tempByte)
			if _, ok := keywords[tempChar]; ok {
				s.setLex(tempChar, keywords[tempChar])
				s.input.Next()
				return nil
			}
			return s.input.Collapse("readOp cannot recognize")
		}
		// endswitch
	default:
		s.input.Next()
		return s.input.Collapse("readOp cannot recognize")
	}
	return nil
}

func main() {
	// program := " begin x := 9; if x > 9 then\r\r\n x := 2 * x + 1 / 3; \n end #"

	program := `
	begin @@ x := 9;
	if x > 9 then
		x @ := 2 * x + 1 / 3
	end
	`

	scanner := NewScanner(NewInput(program))

	token, syn := scanner.Next()
	fmt.Printf("<%s , %d>\n", token, syn)

	for !scanner.EOF() {
		token, syn = scanner.Next()
		fmt.Printf("<'%s', %d>\n", token, syn)
	}
}
