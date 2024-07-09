package main

import (
	"osLab1/algorithm"
)

func main() {
	processes := GenerateProcesses()

	timeSlice := 2
	algorithm.RoundRobin(processes, timeSlice)
	algorithm.ShortestJobFirst(processes)
}
