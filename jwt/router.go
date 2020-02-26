package jwt

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())


	r.GET("/auth", GetAuth)

	login := r.Group("/login")
	login.Use(JWT())
	{
		login.GET("/check_user", getUsrInfo)
	}

	return r
}