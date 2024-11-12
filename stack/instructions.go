package stack

type Instruction int

const (
	NOOP Instruction = iota
	Next
	Recurse
	Stop
)
