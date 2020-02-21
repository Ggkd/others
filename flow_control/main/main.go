package main

import (
	"github.com/Ggkd/flow_control"
	"net/http"
)

func main() {
	r := flow_control.RegisterRouter()
	m := flow_control.NewMiddleWare(r, 3)
	http.ListenAndServe(":8899", m)
}