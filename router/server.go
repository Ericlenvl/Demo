package router

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	RegisterRouters(r)
	//监听端口默认为8080
	r.Run(":8000")
}
