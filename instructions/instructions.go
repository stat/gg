package instructions

type Instruction int

const (
	NOOP Instruction = iota
	Next
	Recurse
	Stop
)
