package router

import (
	"github.com/gin-gonic/gin"
	"hxj/controller"
)

func RegisterRouters(r *gin.Engine) {
	g := r.Group("test")
	{
		g.GET("member", controller.GetMemberById)
	}
}