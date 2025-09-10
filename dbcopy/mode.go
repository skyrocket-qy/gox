package dbcopy

type Mode int

const (
	ModeBasic   Mode = iota // copy if dst data not exist
	ModeAppend              // directly append
	ModeReplace             // replace old data
)

const BatchSize = 500
