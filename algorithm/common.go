package algorithm

import (
	"osLab1/enums"
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

func convertProcessStatus(p *models.Process, newStatus enums.ProcessState, time int) {

	logger.WarningF("<%d ms> [进程 %d] 状态转换：「%s」=> 「%s」",
		time,
		p.PID,
		enums.GetStateName(p.State),
		enums.GetStateName(newStatus))
	p.State = newStatus

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
