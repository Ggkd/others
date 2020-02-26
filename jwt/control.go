package jwt

import (
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")


	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	isExist := CheckAuth(username, password)
	if isExist {
		token, err := GenerateToken(username, password)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token

			code = e.SUCCESS
		}

	} else {
		code = e.ERROR_AUTH
	}


	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

func getUsrInfo(c *gin.Context)  {
	userInfo := auth{
		Username: "luffy",
		Password: "123123",
	}
	c.JSON(200, userInfo)
}