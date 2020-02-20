package main

import (
	"fmt"
	"time"
	"github.com/jasonlvhit/gocron"
)

func task1() {
	fmt.Println("task1 running ", time.Now())
}
func task2() {
	fmt.Println("task2 running ", time.Now())
}

func main() {
	s := gocron.NewScheduler()
	s.Every(1).Seconds().Do(task1)
	s.Every(5).Seconds().Do(task2)

	sc := s.Start() // 获取通道
	go test(s, sc)  // 后台等待运行
	<-sc            // 当通道关闭时，将会执行并退出
}

func test(s *gocron.Scheduler, sc chan bool) {
	time.Sleep(10 * time.Second)
	s.Remove(task1) //移除 task1
	time.Sleep(10 * time.Second)
	s.Clear()
	fmt.Println("All task removed")
	close(sc) // 关闭通道
}