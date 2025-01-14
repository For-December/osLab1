package algorithm

import (
	"osLab1/enums"
	"osLab1/models"
	"osLab1/utls/logger"
)

// MultilevelFeedbackQueue 模拟多级反馈队列调度算法
// 参考：https://blog.csdn.net/weixin_44730681/article/details/109270711
func MultilevelFeedbackQueue(rawProcesses []models.Process, timeSlices []int) {

	calculateProcesses := make([]models.Process, len(rawProcesses))
	// 复制，防止修改外部数组的元素
	copy(calculateProcesses, rawProcesses)

	// 保存原始数据的地址，以便直接修改原始数据（为了最终统计）
	processes := make([]*models.Process, len(calculateProcesses))

	// 数据预处理
	for i := range calculateProcesses {
		calculateProcesses[i].StartTime = -1
		processes[i] = &calculateProcesses[i]
	}

	queues := make([]models.Queue, len(timeSlices))
	time := 0
	for len(processes) > 0 || !allQueuesEmpty(queues) {

		// 所有到达时间为当前时间的进程加入第一个队列
		addNewProcessToQueue(&processes, &queues[0], time)

		// 优先级从高到低，找到第一个非空队列
		var currentQueue *models.Queue
		var queueLevel int
		for i := range queues {
			if !queues[i].IsEmpty() {
				currentQueue = &queues[i]
				queueLevel = i
				break
			}
		}

		// 如果所有队列都为空，CPU空闲，时间流逝
		if currentQueue == nil {
			time++
			continue
		}

		// 第一个进程出队，开始运行
		p := currentQueue.Dequeue()

		// 如果进程第一次运行，设置开始时间为当前时间
		if p.StartTime == -1 {
			p.StartTime = time
		}

		// 就绪 => 运行
		convertProcessStatus(p, enums.Running, time)

		// 运行进程，直到时间片用完或进程完成
		processRunning(p, time)
		runTime := min(timeSlices[queueLevel], p.RemainingTime)
		time += runTime
		p.RemainingTime -= runTime

		// 每次时间流逝都需要将当前时间到达的进程加入队列
		// 在这里立即添加，以实现：如果时间片用完和新进程到达同时发生，认为新进程到达先发生
		addNewProcessToQueue(&processes, currentQueue, time)

		// 如果进程执行完毕，设置它的完成时间、等待时间和响应时间
		if p.RemainingTime == 0 {
			p.FinishTime = time
			p.WaitingTime = p.FinishTime - p.ArrivalTime - p.ExecuteTime
			p.ResponseTime = p.StartTime - p.ArrivalTime

			// 进程完成
			processFinish(p, time)
		} else {
			// 运行 => 就绪
			convertProcessStatus(p, enums.Ready, time)

			// 如果进程未执行完，重新加入队列，并降低优先级
			if queueLevel+1 < len(queues) {
				logger.WarningF("<%d ms> [进程 %d] 进入下一级优先队列",
					time,
					p.PID)
				queues[queueLevel+1].Enqueue(p)
			} else {
				// 如果已经是最低优先级，重新加入之前的队列队尾
				queues[queueLevel].Enqueue(p)
			}
		}
	}

	// 指标计算
	calculateMetrics(calculateProcesses, time)
}

// allQueuesEmpty 检查所有队列是否为空
func allQueuesEmpty(queues []models.Queue) bool {
	for _, q := range queues {
		if !q.IsEmpty() {
			return false
		}
	}
	return true
}
