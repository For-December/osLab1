package algorithm

import (
	"osLab1/enums"
	"osLab1/models"
	"sort"
)

// ShortestJobFirst 短作业优先调度算法
func ShortestJobFirst(processes []models.Process) {
	// 数据预处理
	for i := range processes {
		processes[i].StartTime = -1
	}

	queue := models.Queue{}
	time := 0

	for len(processes) > 0 || !queue.IsEmpty() {

		// 将所有 当前时刻（time） 到达的进程加入队列
		addNewProcessToQueue(&processes, &queue, time)

		// 队列里的进程按照剩余时间排序，剩余时间短的进程优先执行
		sort.Slice(queue.Items, func(i, j int) bool {
			return queue.Items[i].RemainingTime < queue.Items[j].RemainingTime
		})

		// 当前时间没有新进程则时间流逝
		if queue.IsEmpty() {
			time++
			continue
		}

		// 从队列中取出一个进程，开始运行
		p := queue.Dequeue()

		// 如果进程第一次运行，设置开始时间为当前时间
		if p.StartTime == -1 {
			p.StartTime = time
		}

		// 就绪 => 运行
		convertProcessStatus(p, enums.Running, time)

		// 运行进程，直到进程完成
		processRunning(p, time)
		time += p.RemainingTime
		p.RemainingTime = 0

		// 每次时间流逝都需要将当前时间到达的进程加入队列
		// 在这里立即添加，以实现：如果时间片用完和新进程到达同时发生，认为新进程到达先发生
		addNewProcessToQueue(&processes, &queue, time)

		// 设置进程的完成时间、等待时间和响应时间
		p.FinishTime = time
		p.WaitingTime = p.FinishTime - p.ArrivalTime - p.ExecuteTime
		p.ResponseTime = p.StartTime - p.ArrivalTime

		// 进程完成
		processFinish(p, time)
	}
}
