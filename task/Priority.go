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
