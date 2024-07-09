package main

import (
	"osLab1/utls/logger"
)

// RoundRobin 模拟时间片轮转调度算法
// 参考1：https://en.wikipedia.org/wiki/Round-robin_scheduling
// 参考2: https://c.biancheng.net/view/1247.html
// 该算法中，将一个较小时间单元定义为时间量或时间片。时间片的大小通常为 10~100ms。
// 就绪队列作为循环队列。CPU 调度程序循环整个就绪队列，为每个进程分配不超过一个时间片的 CPU。
// 时间单位：ms
func RoundRobin(processes []Process, timeSlice int) {
	// 数据预处理
	for i := range processes {
		processes[i].StartTime = -1
	}

	queue := Queue{}

	time := 1 // 模拟当前时间，判断进程是否到达

	// 轮询队列，直到所有进程完成
	for len(processes) > 0 || !queue.IsEmpty() {

		// 将所有 当前时刻（time） 到达的进程加入队列
		for len(processes) > 0 && processes[0].ArrivalTime <= time {

			logger.DebugF("[ARRIVE]: Process %d arrived at time %d",
				processes[0].PID, time)

			queue.Enqueue(processes[0])
			processes = processes[1:]
		}

		// 如果队列为空，当前时刻没有进程到达，时间流逝，CPU 等待下一个进程到达
		if queue.IsEmpty() {
			time++
			continue
		}

		// 处理器从队列中取出一个进程，开始运行它（调度）
		p := queue.Dequeue()

		// 如果进程第一次运行，设置开始时间为当前时间
		if p.StartTime == -1 {
			p.StartTime = time
		}

		// 运行进程，直到时间片用完或进程完成
		logger.InfoF("[RUNNING]: Process %d is running at time %d",
			p.PID, time)
		runTime := min(timeSlice, p.RemainingTime)
		time += runTime
		p.RemainingTime -= runTime

		// 如果进程执行完毕，设置它的完成时间、等待时间和响应时间
		if p.RemainingTime == 0 {
			p.FinishTime = time
			p.WaitingTime = p.FinishTime - p.ArrivalTime - p.ExecuteTime
			p.ResponseTime = p.StartTime - p.ArrivalTime
			logger.WarningF("[FINISHED]: Process %d finished at time %d",
				p.PID, p.FinishTime)
		} else {
			// 将该进程重新放入队列，等待下一次调度
			queue.Enqueue(*p)
		}
	}
}
