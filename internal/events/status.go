package events

// Status is the lifecycle state of an agent
type Status int

const (
	StatusPending Status = iota
	StatusRunning
	StatusWaiting
	StatusDone
	StatusFailed
)

// String return a readable name such as "running"
func (s Status) String() string {
	switch s {
	case StatusPending:
		return "pending"
	case StatusRunning:
		return "running"
	case StatusWaiting:
		return "waiting"
	case StatusDone:
		return "done"
	case StatusFailed:
		return "failed"
	default:
		return "unknown"
	}
}

// Terminal reports whether the agent has finished for good
func (s Status) Terminal() bool {
	return s == StatusDone || s == StatusFailed
}
