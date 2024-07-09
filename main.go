package main

import (
	"fmt"
	"osLab1/algorithm"
)

func main() {
	processes := GenerateProcesses()

	timeSlice := 2
	algorithm.RoundRobin(processes, timeSlice)
	fmt.Println("############")
	algorithm.ShortestJobFirst(processes)

	fmt.Println("############")
	algorithm.MultilevelFeedbackQueue(processes, []int{1, 2, 4})
}
