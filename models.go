package main

// Process 自定义模拟进程 PCB
type Process struct {
	PID         int // 进程 ID
	ArrivalTime int // 到达时间
	ExecuteTime int // 运行时间
	Priority    int // 优先级

	RemainingTime int // 剩余时间
	FinishTime    int // 完成时间
	StartTime     int // 开始时间
	WaitingTime   int // 等待时间
	ResponseTime  int // 响应时间
}

// Queue 进程队列
type Queue struct {
	items []Process
}

// Enqueue 添加进程到队列
func (q *Queue) Enqueue(p Process) {
	q.items = append(q.items, p)
}

// Dequeue 移除队首进程，并返回该进程
func (q *Queue) Dequeue() *Process {
	if len(q.items) == 0 {
		return nil
	}
	p := q.items[0]
	q.items = q.items[1:]
	return &p
}

// IsEmpty 检查队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Peek 查看队列中的第一个进程，不移除
func (q *Queue) Peek() *Process {
	if len(q.items) == 0 {
		return nil
	}
	return &q.items[0]
}
