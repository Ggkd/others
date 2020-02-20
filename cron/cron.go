package main

import (
	"fmt"
	"github.com/robfig/cron"
)


type Task struct {
	
}

func (t Task) Run() {
	fmt.Println("task running")
}

func main() {
	i := 0
	c := cron.New()
	spec := "*/3 * * * * ?"  //每三秒运行一次
	c.AddFunc(spec, func() {
		i++
		fmt.Println("cron running:", i)
	})

	c.AddJob(spec, Task{})
	c.Start()

	defer c.Stop()

	select {}
}