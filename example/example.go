package main

import (
	"fmt"

	"github.com/insionng/taskcron"
)

func task() {
	fmt.Println("I am runnning task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}

func main() {
	// Do jobs with params
	taskcron.Every(1).Second().Do(taskWithParams, 1, "hello")

	// Do jobs without params
	taskcron.Every(1).Second().Do(task)
	taskcron.Every(2).Seconds().Do(task)
	taskcron.Every(1).Minute().Do(task)
	taskcron.Every(2).Minutes().Do(task)
	taskcron.Every(1).Hour().Do(task)
	taskcron.Every(2).Hours().Do(task)
	taskcron.Every(1).Day().Do(task)
	taskcron.Every(2).Days().Do(task)

	// Do jobs on specific weekday
	taskcron.Every(1).Monday().Do(task)
	taskcron.Every(1).Thursday().Do(task)

	// function At() take a string like 'hour:min'
	taskcron.Every(1).Day().At("10:30").Do(task)
	taskcron.Every(1).Monday().At("18:30").Do(task)

	// remove, clear and next_run
	_, time := taskcron.NextRun()
	fmt.Println(time)

	// taskcron.Remove(task)
	// taskcron.Clear()

	// function Start start all the pending jobs
	<-taskcron.Start()

	// also , you can create a your new scheduler,
	// to run two scheduler concurrently
	s := taskcron.NewScheduler()
	s.Every(3).Seconds().Do(task)
	<-s.Start()
}
