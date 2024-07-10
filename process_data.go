package main

import "osLab1/models"

// GenerateProcesses 生成测试进程
// 该测试进程来自王道：https://blog.csdn.net/yrx420909/article/details/104363553
func GenerateProcesses() []models.Process {
	processes := []models.Process{
		{PID: 1, ArrivalTime: 0, ExecuteTime: 5, RemainingTime: 5, Priority: 2},
		{PID: 2, ArrivalTime: 2, ExecuteTime: 4, RemainingTime: 4, Priority: 1},
		{PID: 3, ArrivalTime: 4, ExecuteTime: 1, RemainingTime: 1, Priority: 3},
		{PID: 4, ArrivalTime: 5, ExecuteTime: 6, RemainingTime: 6, Priority: 2},
	}
	return processes
}
