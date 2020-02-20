package shared

//go:generate stringer -type=ChainHandledState

type ChainHandledState int

func (p ChainHandledState) Failed() bool {
	return p == ChainHandledStateThenFailed ||
		p == ChainHandledStateElseFailed
}

const (
	ChainHandledStateThenFailed ChainHandledState = iota
	ChainHandledStateElseFailed
	ChainHandledStateUnhandled
	ChainHandledStateThen
	ChainHandledStateElse
)
