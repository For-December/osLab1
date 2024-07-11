package algorithm

import (
	"fmt"
	"osLab1/enums"
	"osLab1/models"
	"osLab1/utls/logger"
)

// 两个内部数据地址传递
func addNewProcessToQueue(processes *[]*models.Process, queue *models.Queue, time int) {
	// 将所有 当前时刻（time） 到达的进程加入队列
	for len(*processes) > 0 && (*processes)[0].ArrivalTime <= time {

		logger.DebugF("<%d ms> [进程 %d] 到达",
			time,
			(*processes)[0].PID)

		queue.Enqueue((*processes)[0])
		*processes = (*processes)[1:]
	}
}

func convertProcessStatus(p *models.Process, newStatus enums.ProcessState, time int) {

	logger.WarningF("<%d ms> [进程 %d] 状态转换：「%s」=> 「%s」",
		time,
		p.PID,
		enums.GetStateName(p.State),
		enums.GetStateName(newStatus))
	p.State = newStatus

}

// 计算各项性能指标
func calculateMetrics(processes []models.Process, totalTime int) {
	// 总周转时间、总带权周转时间、总等待时间、总响应时间
	var totalTurnaroundTime, totalWeightedTurnaroundTime, totalWaitingTime, totalResponseTime int
	for _, p := range processes {
		turnaroundTime := p.FinishTime - p.ArrivalTime
		weightedTurnaroundTime := float64(turnaroundTime) / float64(p.ExecuteTime)
		totalTurnaroundTime += turnaroundTime
		totalWeightedTurnaroundTime += int(weightedTurnaroundTime)
		totalWaitingTime += p.WaitingTime
		totalResponseTime += p.ResponseTime
	}

	n := len(processes)
	averageTurnaroundTime := float64(totalTurnaroundTime) / float64(n)
	averageWeightedTurnaroundTime := float64(totalWeightedTurnaroundTime) / float64(n)
	averageWaitingTime := float64(totalWaitingTime) / float64(n)
	averageResponseTime := float64(totalResponseTime) / float64(n)
	cpuUtilization := float64(totalTime) / float64(totalTime) * 100

	fmt.Printf("平均周转时间: %.2f\n", averageTurnaroundTime)
	fmt.Printf("平均带权周转时间: %.2f\n", averageWeightedTurnaroundTime)
	fmt.Printf("平均等待时间: %.2f\n", averageWaitingTime)
	fmt.Printf("平均响应时间: %.2f\n", averageResponseTime)
	fmt.Printf("CPU 利用率: %.2f%%\n", cpuUtilization)
}

func processRunning(p *models.Process, time int) {
	logger.InfoF("<%d ms>  [进程 %d] 运行",
		time,
		p.PID)
}

func processFinish(p *models.Process, time int) {
	logger.ErrorF("<%d ms> [进程 %d] 完成",
		time,
		p.PID)
}
