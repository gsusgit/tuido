package storage

// PriorityOrder returns sort rank for priorities (lower = higher priority).
func PriorityOrder(p Priority) int {
	switch p {
	case PriorityHigh:
		return 0
	case PriorityMedium:
		return 1
	default:
		return 2
	}
}
