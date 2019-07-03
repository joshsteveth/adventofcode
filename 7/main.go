package main

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"time"

	"github.com/joshsteveth/adventofcode/util"
)

func initList() (
	todoList map[string][]string,
	allTask map[string]struct{},
	remainingTask map[string]struct{},
) {

	todoList = map[string][]string{}
	allTask = map[string]struct{}{}
	remainingTask = map[string]struct{}{}

	lines, err := util.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`Step ([A-Z]) must be finished before step ([A-Z]) can begin.`)

	for _, l := range lines {
		substr := re.FindStringSubmatch(l)
		mustFinish, todo := substr[1], substr[2]

		// Step C must be finished before step A can begin.
		// todoList[A] = [C]
		todoList[todo] = append(todoList[todo], mustFinish)

		// register existing tasks
		allTask[todo] = struct{}{}
		allTask[mustFinish] = struct{}{}

	}

	// populate remaining task as well
	for t, _ := range allTask {
		remainingTask[t] = struct{}{}
	}

	return
}

func main() {

	timeout := time.Second * 3
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ch := make(chan int)

	tasks := []func(chan<- int){
		taskOne,
		taskTwo,
	}

	for _, t := range tasks {
		go t(ch)
	}

	for i := 0; i < len(tasks); i++ {
		select {
		case <-ctx.Done():
			panic("context cancelled")
		case i := <-ch:
			fmt.Printf("Task#%d done! \n", i)
		}

	}

}

func solveTask(
	task string,
	todoList map[string][]string,
	remainingTask map[string]struct{},
) {
	delete(remainingTask, task)

	removeTaskFromList := func(list []string) []string {
		newList := []string{}
		for _, l := range list {
			if l != task {
				newList = append(newList, l)
			}
		}

		return newList
	}

	for key, list := range todoList {
		newList := removeTaskFromList(list)

		if len(newList) == 0 {
			delete(todoList, key)
			continue
		}

		todoList[key] = newList
	}
}

func taskOne(ch chan<- int) {

	todoList, allTask, remainingTask := initList()

	var result string

	//for i := 0; i < 2; i++ {
	for len(remainingTask) != 0 {

		// let's get tasks that can be solved directly
		solveableTask := []string{}

		for task, _ := range allTask {
			if _, ok := todoList[task]; !ok {

				if _, stillPending := remainingTask[task]; !stillPending {
					continue
				}
				solveableTask = append(solveableTask, task)
			}
		}

		sort.Strings(solveableTask)

		result += solveableTask[0]
		solveTask(solveableTask[0], todoList, remainingTask)

	}

	fmt.Printf("task one result: %s\n", result)

	ch <- 1
}

func calculateAmount(task string) int {
	return int(task[0]) - 64
}

const (
	numWorkers = 5
	baseTime   = 60
)

var timePassed = 0

type worker struct {
	isFree        bool
	workingOnTask string
	remainingTime int
}

func (w *worker) assignTask(task string) {
	w.isFree = false
	w.workingOnTask = task
	w.remainingTime = baseTime + calculateAmount(task)
}

func taskTwo(ch chan<- int) {

	todoList, allTask, remainingTask := initList()

	workers := []*worker{}
	assignedTask := map[string]struct{}{}

	// initiates x number of workers
	for i := 0; i < numWorkers; i++ {
		workers = append(workers, &worker{
			isFree: true,
		})
	}

	assignTaskToWorker := func(task string, workers []*worker) {

		if _, ok := assignedTask[task]; ok {
			return
		}

		for i := 0; i < numWorkers; i++ {

			w := workers[i]

			if !w.isFree {
				continue
			}

			// if worker is free then let's assign this task to him
			w.assignTask(task)
			assignedTask[task] = struct{}{}
			return

		}
	}

	for len(remainingTask) != 0 {

		// let's get tasks that can be solved directly
		solveableTask := []string{}

		for task, _ := range allTask {
			if _, ok := todoList[task]; !ok {

				if _, stillPending := remainingTask[task]; !stillPending {
					continue
				}
				solveableTask = append(solveableTask, task)
			}
		}

		sort.Strings(solveableTask)

		// assign the task to each workers
		for _, task := range solveableTask {
			assignTaskToWorker(task, workers)
		}

		//fmt.Printf("time passed %d worker1 %v worker2 %v worker3 %v worker4 %v worker5 %v\n", timePassed, *workers[0], *workers[1], *workers[2], *workers[3], *workers[4])

		// let's time jump to the next phase
		// that is when next worker finishes his job

		var busyWorkers []*worker
		for _, worker := range workers {
			if !worker.isFree {
				busyWorkers = append(busyWorkers, worker)
			}
		}

		sort.Slice(busyWorkers, func(i, j int) bool { return busyWorkers[i].remainingTime < busyWorkers[j].remainingTime })

		nextTaskFinish := busyWorkers[0].remainingTime
		solveTask(busyWorkers[0].workingOnTask, todoList, remainingTask)

		timePassed += nextTaskFinish

		for _, worker := range busyWorkers {
			worker.remainingTime = worker.remainingTime - nextTaskFinish
			if worker.remainingTime <= 0 {
				delete(assignedTask, worker.workingOnTask)
				worker.isFree = true
				worker.workingOnTask = ""
			}
		}
	}

	fmt.Printf("time passed: %ds\n", timePassed)

	ch <- 2
}
