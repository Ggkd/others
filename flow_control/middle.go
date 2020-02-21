package flow_control

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type MiddleWare struct {
	router *httprouter.Router
	cl *ConnLimit
}

func NewMiddleWare(r *httprouter.Router, limit int) http.Handler {
	return MiddleWare{
		router: r,
		cl:    	NewConnLimit(limit),
	}
}

func (m MiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.cl.getConn(){
		fmt.Println("请求量太多")
		return
	}
	m.router.ServeHTTP(w, r)
	m.cl.ReleaseConn()
}

func RegisterRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/test", test)
	return router
}

func test(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	//为了演示效果这块设置了等待
	time.Sleep(time.Second * 5)
	w.Write([]byte("get success"))
}