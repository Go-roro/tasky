package task

type Priority int

const (
	PriorityLow    = 1
	PriorityMedium = 2
	PriorityHigh   = 3
)

func (p Priority) ValidPriority() bool {
	return p >= PriorityLow && p <= PriorityHigh
}

func (p Priority) Symbol() string {
	switch p {
	case PriorityLow:
		return "🟢"
	case PriorityMedium:
		return "🟠"
	case PriorityHigh:
		return "🔴"
	default:
		return "🟠"
	}
}
