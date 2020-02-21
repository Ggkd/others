package flow_control

import "fmt"

type ConnLimit struct {
	currentConn 	int
	bucket         chan int
}

func NewConnLimit(limit int) *ConnLimit {
	return &ConnLimit{
		currentConn: 0,
		bucket:      make(chan int, limit),
	}
}

func (cl *ConnLimit) getConn() bool {
	if len(cl.bucket) > cl.currentConn {
		fmt.Println("连接池已满")
		return false
	}
	cl.bucket <- 11
	return true
}

func (cl *ConnLimit) ReleaseConn() {
	<- cl.bucket
}