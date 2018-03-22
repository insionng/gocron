## taskcron: A Golang Job Scheduling Package.

taskcron is a Golang job scheduling package which lets you run Go functions periodically at pre-determined interval using a simple, human-friendly syntax.

taskcron is a Golang implementation of Ruby module [clockwork](<https://github.com/tomykaira/clockwork>) and Python job scheduling package [schedule](<https://github.com/dbader/schedule>), and personally, this package is my first Golang program, just for fun and practice.

See also this two great articles:
* [Rethinking Cron](http://adam.heroku.com/past/2010/4/13/rethinking_cron/)
* [Replace Cron with Clockwork](http://adam.heroku.com/past/2010/6/30/replace_cron_with_clockwork/)

Back to this package, you could just use this simple API as below, to run a cron scheduler.

``` go
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

	taskcron.Remove(task)
	taskcron.Clear()

	// function Start start all the pending jobs
	<- taskcron.Start()

	// also , you can create a your new scheduler,
	// to run two scheduler concurrently
	s := taskcron.NewScheduler()
	s.Every(3).Seconds().Do(task)
	<- s.Start()

}
```
and full test cases and [document](http://godoc.org/github.com/insionng/taskcron) will be coming soon.

Once again, thanks to the great works of Ruby clockwork and Python schedule package. BSD license is used, see the file License for detail.

Hava fun!
