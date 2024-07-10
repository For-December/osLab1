package algorithm

import (
	"fmt"
	"osLab1/models"
)

// MultilevelFeedbackQueue 模拟多级反馈队列调度算法
// 参考：https://blog.csdn.net/weixin_44730681/article/details/109270711
func MultilevelFeedbackQueue(processes []models.Process, timeSlices []int) {
	queues := make([]models.Queue, len(timeSlices))
	time := 0
	for len(processes) > 0 || !allQueuesEmpty(queues) {

		// 所有到达时间为当前时间的进程加入第一个队列
		for len(processes) > 0 && processes[0].ArrivalTime <= time {

			queues[0].Enqueue(processes[0])
			processes = processes[1:]
		}

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

		// 运行进程，直到时间片用完或进程完成
		runTime := min(timeSlices[queueLevel], p.RemainingTime)
		time += runTime
		p.RemainingTime -= runTime

		// 如果进程执行完毕，设置它的完成时间、等待时间和响应时间
		if p.RemainingTime == 0 {
			p.FinishTime = time
			p.WaitingTime = p.FinishTime - p.ArrivalTime - p.ExecuteTime
			p.ResponseTime = p.StartTime - p.ArrivalTime
			fmt.Printf("Process %d finished at time %d\n", p.PID, p.FinishTime)
		} else {
			// 如果进程未执行完，重新加入队列，并降低优先级
			if queueLevel+1 < len(queues) {
				queues[queueLevel+1].Enqueue(*p)
			} else {
				// 如果已经是最低优先级，重新加入之前的队列队尾
				queues[queueLevel].Enqueue(*p)
			}
		}
	}
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