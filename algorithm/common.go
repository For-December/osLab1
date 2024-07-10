package algorithm

import (
	"osLab1/models"
	"osLab1/utls/logger"
)

// 两个内部数据地址传递
func addNewProcessToQueue(processes *[]models.Process, queue *models.Queue, time int) {
	// 将所有 当前时刻（time） 到达的进程加入队列
	for len(*processes) > 0 && (*processes)[0].ArrivalTime <= time {

		logger.DebugF("<%d ms> [进程 %d] 到达",
			time,
			(*processes)[0].PID)

		queue.Enqueue((*processes)[0])
		*processes = (*processes)[1:]
	}
}
