package main

import "github.com/Ggkd/jwt"

func main() {
	app := jwt.InitRouter()
	app.Run("127.0.0.1:8989")
}