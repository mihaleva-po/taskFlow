package common

const (
	StatusToDo       = "TODO"
	StatusInProgress = "IN PROGRESS"
	StatusDone       = "DONE"
)

func IsValidStatus(status string) bool {
	switch status {
	case StatusDone, StatusInProgress, StatusToDo:
		return true

	default:
		return false
	}
}
