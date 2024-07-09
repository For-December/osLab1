package models

import "osLab1/enums"

// Process 自定义模拟进程 PCB
type Process struct {
	PID         int // 进程 ID
	ArrivalTime int // 到达时间
	ExecuteTime int // 运行时间
	Priority    int // 优先级

	State enums.ProcessState // 当前的进程状态

	RemainingTime int // 剩余时间
	FinishTime    int // 完成时间
	StartTime     int // 开始时间
	WaitingTime   int // 等待时间
	ResponseTime  int // 响应时间
}
