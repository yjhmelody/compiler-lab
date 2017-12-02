package lexer

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

// SkipWhitespace will skip whitespace
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
	var err error
	// fmt.Println("==========", string(s.input.Peek()))
	s.SkipWhitespace()
	if ch := s.input.Peek(); ch == '#' {
		s.setLex("#", SHARP)
		s.input.Next()
	} else if IsDigit(ch) {
		err = s.readNum()
	} else if IsLetter(ch) {
		err = s.readID()
	} else if IsOpChar(ch) {
		err = s.readOp()
	} else {
		// fmt.Println("==========", string(s.input.Peek()))
		fmt.Println("read 不合法的字符串" + string(ch))
		s.input.Next()
		s.read()
		return
	}
	if err != nil {
		// fmt.Println("234242424242", string(s.input.Peek()))
		fmt.Println(err)
		s.read()
	}
}

// readNum read the num type
func (s *Scanner) readNum() error {
	// first digit [1-9]
	ch := s.input.Peek()
	str := string(ch)

	if ch == '0' {
		// if ch == '0'但下一个字符为字母则跳过并且报错
		if ch, _ := s.input.Next(); IsLetter(ch) || IsDigit(ch) {
			for ch := s.input.Peek(); IsLetter(ch) || IsDigit(ch); {
				str += string(ch)
				ch, _ = s.input.Next()
			}
			return s.input.Collapse("readNum 不合法的字符串:" + str)
		}

		s.setLex(str, NUM)
	} else {
		// digit [0-9]
		// 一直读完数字
		for {
			ch, _ = s.input.Next()
			if IsDigit(ch) {
				str += string(ch)
			} else {
				break
			}
		}
		// 数字后面紧接着字母则报错并且移动到第一个运算符号或空白符
		if IsLetter(ch) {
			for IsLetter(ch) || IsDigit(ch) {
				str += string(ch)
				ch, _ = s.input.Next()
				// fmt.Println("343434", str)
			}
			// fmt.Println("2323223", string(s.input.Peek()))
			return s.input.Collapse("readNum 不合法的字符串:" + str)
		}
		s.setLex(str, NUM)
	}
	return nil
}

// readID read the identifier and keywords
func (s *Scanner) readID() error {
	// first letter [A-Za-z]
	ch := s.input.Peek()
	str := string(ch)
	// letter [A-Za-z0-9]
	for {
		ch, _ = s.input.Next()
		if IsLetter(ch) || IsDigit(ch) {
			str += string(ch)
		} else {
			// fmt.Println("ID:"+str, string(ch))
			break
		}
	}
	// recognize keywords
	if kw, ok := keywords[str]; ok {
		s.setLex(str, kw)
	} else {
		s.setLex(str, ID)
	}
	return nil
}

// readOp read the operations
func (s *Scanner) readOp() error {
	// operator: :  :=  +  -  *  /  <  <=  <>  >  >=  =  ;  (  )  #
	var str string
	// fmt.Println("2323223", string(s.input.Peek()))
	switch ch := s.input.Peek(); ch {
	case '=', '+', '-', '*', '/', ';', '(', ')', '#':
		str += string(ch)
		s.setLex(str, keywords[str])
		s.input.Next()
		return nil
	case ':', '<', '>':
		str += string(ch)

		// switch
		switch ch, _ := s.input.Next(); ch {
		case '=', '>':
			str += string(ch)
			if _, ok := keywords[str]; ok {
				s.setLex(str, keywords[str])
				s.input.Next()
				return nil
			}
		default:
			if _, ok := keywords[str]; ok {
				s.setLex(str, keywords[str])
				// s.input.Next()
				return nil
			}
			return s.input.Collapse("readOp 不合法的字符串")
		}
		// endswitch
	default:
		s.input.Next()
		return s.input.Collapse("readOp 不合法的字符串")
	}
	return nil
}

// LexParse write syn to a file
// func (s *Scanner) LexParse() {
// 	token, syn := s.Next()
// 	fmt.Printf("%s,%d>\n", token, syn)
// 	for !s.EOF() {
// 		token, syn = s.Next()
// 		fmt.Printf("%s,%d\n", token, syn)
// 	}
// }

// func main() {

// 	program := `9x9x
// 	0099
// 	??$$
// 	++
// 	begin 9x:=?$00999; if x%><<>9 t99he&n x:=2**x+1/3; end #
// 	`

// 	scanner := NewScanner(NewInput(program))
// 	token, syn := scanner.Next()
// 	fmt.Printf("<'%s', %d>\n", token, syn)
// 	for !scanner.EOF() {
// 		token, syn = scanner.Next()
// 		fmt.Printf("<'%s', %d>\n", token, syn)
// 	}
// }
