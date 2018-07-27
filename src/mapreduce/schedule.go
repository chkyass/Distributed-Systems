package mapreduce

import (
	"fmt"
	"sync"

	"github.com/golang-collections/collections/stack"
)

var wg = sync.WaitGroup{}

func execTask(address string, taskChan chan DoTaskArgs, failsChan chan DoTaskArgs) {
	// Wait for tasks from the scheduler
	for {
		select {
		case task := <-taskChan:
			if !call(address, "Worker.DoTask", task, nil) {
				failsChan <- task
			} else {
				wg.Done()
			}
		}
	}
}

func schedule(jobName string, mapFiles []string, nReduce int, phase jobPhase, registerChan chan string) {
	var ntasks int
	var n_other int
	var failsChan = make(chan DoTaskArgs)

	// To make select and WaitGroup compatible
	quit := make(chan bool)
	// Contains the DoTaskArgs for rpc calls
	tasks := stack.Stack{}

	switch phase {
	case mapPhase:
		ntasks = len(mapFiles)
		n_other = nReduce
	case reducePhase:
		ntasks = nReduce
		n_other = len(mapFiles)
	}

	// Used to send tasks to routines
	// length set to ntask+1 to not block in select default part
	var taskChan = make(chan DoTaskArgs, ntasks+1)

	for j := 0; j < ntasks; j++ {
		tasks.Push(DoTaskArgs{jobName, mapFiles[j], phase, j, n_other})
	}

	wg.Add(ntasks)
	go func() {
		wg.Wait()
		quit <- true
	}()

Loop:
	for {
		select {
		case newAddr := <-registerChan:
			go execTask(newAddr, taskChan, failsChan)
			if tasks.Len() > 0 {
				taskChan <- tasks.Pop().(DoTaskArgs)
			}
		case task := <-failsChan:
			tasks.Push(task)
		case <-quit:
			break Loop
		// block if taskChan size isn't enough big
		default:
			if tasks.Len() > 0 {
				taskChan <- tasks.Pop().(DoTaskArgs)
			}
		}
	}

	fmt.Printf("Schedule: %v %v tasks (%d I/Os)\n", ntasks, phase, n_other)

	fmt.Printf("Schedule: %v done\n", phase)
}
