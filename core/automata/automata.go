package automata

// Automata interface defines the required behaviour for something
// to be considered an Automata
type Automata interface {
	Append(other Automata)
	Union(other Automata)
	Loop()
}
