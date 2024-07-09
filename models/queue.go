package models

// Queue 进程队列
type Queue struct {
	Items []Process
}

// Enqueue 添加进程到队列
func (q *Queue) Enqueue(p Process) {
	q.Items = append(q.Items, p)
}

// Dequeue 移除队首进程，并返回该进程
func (q *Queue) Dequeue() *Process {
	if len(q.Items) == 0 {
		return nil
	}
	p := q.Items[0]
	q.Items = q.Items[1:]
	return &p
}

// IsEmpty 检查队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}

// Peek 查看队列中的第一个进程，不移除
func (q *Queue) Peek() *Process {
	if len(q.Items) == 0 {
		return nil
	}
	return &q.Items[0]
}
