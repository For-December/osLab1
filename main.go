package main

import "fmt"

func main() {
	// Generate processes
	processes := GenerateProcesses()

	// Run Round Robin scheduling
	timeSlice := 2
	RoundRobin(processes, timeSlice)
}

func init() {

	// Generate processes
	processes := GenerateProcesses()

	// Print generated processes
	for _, p := range processes {
		fmt.Printf("Process ID: %d, Arrival Time: %d, Burst Time: %d, Priority: %d\n",
			p.PID, p.ArrivalTime, p.ExecuteTime, p.Priority)
	}

	// Example usage
	p1 := Process{PID: 1, ArrivalTime: 0, ExecuteTime: 5, RemainingTime: 5, Priority: 1}
	p2 := Process{PID: 2, ArrivalTime: 1, ExecuteTime: 3, RemainingTime: 3, Priority: 2}

	queue := Queue{}
	queue.Enqueue(p1)
	queue.Enqueue(p2)

	fmt.Println("First process in queue:", queue.Peek())
	dequeuedProcess := queue.Dequeue()
	fmt.Println("Dequeued process:", dequeuedProcess)
	fmt.Println("Queue is empty:", queue.IsEmpty())
}
