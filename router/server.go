package router

import (
	"github.com/gin-gonic/gin"
)

type Testobj struct {
	Rhandler RestHandler
}

func (t *Testobj) Start() {
	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	t.Rhandler.RegisterRouters(r)
	//监听端口默认为8080
	r.Run(":8080")
}
