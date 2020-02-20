package shared

//go:generate stringer -type=ChainState

type ChainState int

func (p ChainState) Failed() bool {
	return p == ChainStateThenFailed ||
		p == ChainStateElseFailed
}

func (p ChainState) Ok() bool {
	return p == ChainStateFinished
}

const (
	ChainStateThenFailed ChainState = iota
	ChainStateElseFailed
	ChainStateUnhandled
	ChainStateFinished
)
