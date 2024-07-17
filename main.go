package main

import (
	"fmt"
	"osLab1/algorithm"
)

func main() {
	processes := GenerateProcesses()

	timeSlice := 2
	fmt.Println("############ 时间片轮转 ############")
	algorithm.RoundRobin(processes, timeSlice)
	fmt.Println("############ 短作业优先 ############")
	algorithm.ShortestJobFirst(processes)

	fmt.Println("############ 多级反馈队列 ############")
	algorithm.MultilevelFeedbackQueue(processes, []int{1, 2, 4})
}
