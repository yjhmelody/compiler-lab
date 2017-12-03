package DFA

import (
	"bytes"
	"fmt"
)

// State store the string as a state id
type State string

// Letter store the string as a letter id
type Letter string

func (s State) String() string {
	return string(s)
}
func (l Letter) String() string {
	return string(l)
}

// DFA store the dfa struct
type DFA struct {
	q      map[State]bool                     // States
	e      map[Letter]bool                    // Alphabet
	d      map[domainelement]*codomainelement // Transition
	q0     State                              // Start State
	f      map[State]bool                     // Terminal States
	done   chan laststate                     // Terminal channel
	input  *Letter                            // Inputs to the DFA
	stop   chan struct{}                      // Stops the DFA
	logger func(State)                        // looger for transitions
}

type domainelement struct {
	l Letter
	s State
}

type codomainelement struct {
	s    State
	exec interface{}
}

type laststate struct {
	s        State
	accepted bool
}

// New returns the dfa
func New() *DFA {
	return &DFA{
		q:      make(map[State]bool),
		e:      make(map[Letter]bool),
		d:      make(map[domainelement]*codomainelement),
		done:   make(chan laststate, 1),
		stop:   make(chan struct{}),
		logger: func(State) {},
	}
}

// SetTransition, argument 'exec' must be a function
// that will supply the next letter if the 'to' state is non-terminal.

func (dfa *DFA) SetTransition(from State, input Letter, to State, exec interface{}) {
	if exec == nil {
		panic("")
	}
	if from == State("") || to == State("") {
		panic("")
	}

	switch exec.(type) {
	case func():
		// f[to] is not a terminal state
		if !dfa.f[to] {
			panic(fmt.Sprintf(""))
		}

	case func() Letter:
		// f[to] is a terminal state
		if dfa.f[to] {
			panic(fmt.Sprintf(""))
		}
	default:
		panic("")
	}

	dfa.q[to] = true
	dfa.q[from] = true
	dfa.e[input] = true

	de := domainelement{l: input, s: from}
	if _, ok := dfa.d[de]; !ok {
		dfa.d[de] = &codomainelement{s: to, exec: exec}
	}
}

// SetStartState just can set one state
func (dfa *DFA) SetStartState(q0 State) {
	dfa.q0 = q0
}

// SetTerminalStates sets some terminal states
func (dfa *DFA) SetTerminalStates(f ...State) {
	for _, q := range f {
		dfa.f[q] = true
	}
}

// SetTransitionLogger set a logger for dfa
func (dfa *DFA) SetTransitionLogger(logger func(State)) {
	dfa.logger = logger
}

// States returns the DFA's all states
func (dfa *DFA) States() []State {
	q := make([]State, len(dfa.q))
	for s := range dfa.q {
		q = append(q, s)
	}
	return q
}

// Alphabet returns the DFA's alphabet
func (dfa *DFA) Alphabet() []Letter {
	e := make([]Letter, len(dfa.e))
	for l := range dfa.e {
		e = append(e, l)
	}
	return e
}

func (dfa *DFA) Run(init interface{}) (State, bool) {
	// assert something
	if init == nil {
		panic("")
	}
	if dfa.q0 == State("") {
		panic("")
	}
	if len(dfa.f) == 0 {
		panic("")
	}
	if _, ok := dfa.q[dfa.q0]; !ok {
		panic("")
	}
	for s := range dfa.f {
		if _, ok := dfa.q[s]; !ok {
			panic("")
		}
	}
	// end assert

	// run the DFA
	go func() {
		defer close(dfa.done)
		// starts at q0
		s := m.q0
		if dfa.f[s] {
			dfa.done <- laststate{s, true}
			return
		} else {
			// otherwise continue reading generated input
			// by starting the next stateful computation
			switch init := init.(type) {
			case func():
				dfa.logger(s)
				init()
			case func() Letter:
				dfa.logger(s)
				l := init()
				dfa.input = &l
			}
		}

		for {
			var stopNow bool
			select {
			case <-dfa.stop:
				stopNow = true
			default:
			}

			if stopNow {
				break
			}

			if dfa.input != nil {
				l := *dfa.input
				// reject upfront if letter is not in alphabet
				if !dfa.e[l] {
					panic("")
				}

				de := domainelement{l: l, s: s}

				if coe := dfa.d[de]; coe != nil {
					s = coe.s
					switch exec := coe.exec.(type) {
					case func():
						dfa.logger(s)
						exec()
					case func():
						dfa.logger(s)
						l := exec()
						dfa.input = &l
					}

					if dfa.f[s] {
						// if the new state is aterminal state
						// the the DFA has accepted the input sequence
						// and it can stop
						dfa.done <- laststate{s, true}
						return
					}
				} else {
					// otherwise stop the dfa with a rejected state,
					// the DFA has rejected the input sequence
					panic("")
				}
			}
		}

		// the caller has cosed the input channel, check if the
		// current state is accepted or rejected by the DFA
		if dfa.f[s] {
			dfa.done <- laststate{s, true}
		} else {
			dfa.done <- laststate{s, false}
		}
	}()

	return dfa.result()
}

func (dfa *DFA) result() (State, bool) {
	t := <-dfa.done
	return t.s, t.accepted
}

// Stop the DFA
func (dfa *DFA) Stop() {
	close(dfa.stop)
}

// GraphViz representation string which can be copy-n-pasted into
// any online tool like http://graphs.grevian.org/graph to get
// a diagram of the DFA.
func (m *DFA) GraphViz() string {
	var buf bytes.Buffer
	buf.WriteString("digraph {\n")
	for do, cdo := range m.d {
		if do.s == m.q0 {
			buf.WriteString(fmt.Sprintf("    \"%s\" -> \"%s\"[label=\"%s\"];\n", do.s, cdo.s, do.l))
		} else if m.f[cdo.s] {
			buf.WriteString(fmt.Sprintf("    \"%s\" -> \"%s\"[label=\"%s\"];\n", do.s, cdo.s, do.l))
		} else {
			buf.WriteString(fmt.Sprintf("    \"%s\" -> \"%s\"[label=\"%s\"];\n", do.s, cdo.s, do.l))
		}
	}
	buf.WriteString("}")
	return buf.String()
}

func (m *DFA) result() (State, bool) {
	t := <-m.done
	return t.s, t.accepted
}
