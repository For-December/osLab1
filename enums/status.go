package enums

type ProcessState int

// 默认初始是就绪态
// 被调度运行变为运行态
// 时间片用完回到就绪态
// 参考：https://cloud.tencent.com/developer/article/2121046
const (
	Ready ProcessState = iota
	Running
	Blocked
)

func GetStateName(state ProcessState) string {
	switch state {
	case Ready:
		return "就绪态"
	case Running:
		return "运行态"
	case Blocked:
		return "阻塞态"
	default:
		return "Unknown"
	}
}
