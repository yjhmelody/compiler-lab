package NFA

// Edge has direction
type Edge byte

// State has not only one edge
type State int

// Transport
type Transport struct {
	state State
	edge  Edge
}

// NFA
type NFA struct {
	// states map[State]bool
	// edges  map[Edge]bool
	// transports map[Transport]State
	state State
	edge  Edge
	// epsilon transport
	next []*NFA
}

var stateCount State
var edgeCount Edge
var 

func parse(str string) *NFA {
	nfa := &NFA{}
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '(':
			j := i + 1
			for j < len(str) {
				if str[j] == ')' {
					break
				}
				j++
			}
			subStr := str[i+1 : j]
			subnfa := parse(subStr)

		// case ')':

		case '*':

		case '|':

		default:
			nfa.edge = Edge(str[i])
			nfa.state = stateCount
			nfa.next
			stateCount++
		}
	}
	return nfa
}

func parseAtom(str string) {

}
