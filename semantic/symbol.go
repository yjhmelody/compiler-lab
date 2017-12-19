package main

import (
	// "github.com/yjhmelody/compiler-lab/lexer"
	"github.com/yjhmelody/compiler-lab/lexer"
)

// Symbol stores symbol's info
type Symbol struct {
	typeName string
	offset   int
}

// SymbolTable manages all symbols
type SymbolTable struct {
	symbols map[string]Symbol
	width   int
}

// mktable(previous)创建一张新的符号表，并返回指向新表的指针。
// 参数previous指向先前创建的符号，放在新符号表的表头。
func mktable(name string, s *Symbol) *SymbolTable {
	st := &SymbolTable{}
	if name == "" {
		return st
	}
	st.symbols[name] = *s
	return st
}

// enter(name, typeName, offset)在table指向的符号表中为名字name建立新表项，
// 同时将类型typeName及相对地址offset放入该表项的属性域中。
func (st *SymbolTable) enter(name, typeName string, offset int) {
	st.symbols[name] = Symbol{typeName, offset}
}

// addwidth(width)将table指向的符号表中
// 所有表项的宽度之和记录在与符号表关联的表头中。
func (st *SymbolTable) addWidth(width int) {
	st.width = width
}

// enterproc(name, newtable)为过程name建立一个新表项，
// 参数newtable指向过程name的符号表。
func (st *SymbolTable) enterproc(name string, s Symbol) {
	st.symbols[name] = s
}

// defined for syntax
const (
	// 29
	EPISILON lexer.Token = lexer.EPISILON + iota
	P
	M
	D
	D2
	T
)
