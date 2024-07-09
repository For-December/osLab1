package enums

type ProcessState int

const (
	Ready ProcessState = iota
	Running
	Blocked
)

func GetStateName(state ProcessState) string {
	switch state {
	case Ready:
		return "Ready"
	case Running:
		return "Running"
	case Blocked:
		return "Blocked"
	default:
		return "Unknown"
	}
}
