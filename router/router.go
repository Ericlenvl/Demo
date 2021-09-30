package router

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"hxj/interfaces"
	"hxj/service"
	"net/http"
	"sync"
)

type RestHandler interface {
	RegisterRouters(engine *gin.Engine)
}

var (
	rOnce sync.Once
	r     RestHandler
)

type resthandler struct {
	userop  interfaces.UserService
}

func NewRestHandler() RestHandler {
	rOnce.Do(func() {
		r = &resthandler{
			userop: service.NewUserService(),
		}
	})
	return r
}

func (h *resthandler) RegisterRouters(r *gin.Engine) {
	g := r.Group("test")
	{
		g.GET("user/:id", h.GetAllUserInfos)
	}
}

func (h *resthandler) GetAllUserInfos(c *gin.Context) {
	//获取请求参数
	id := c.Param("id")
	v, err := h.userop.GetUserInfo(id)
	if err != nil || v != nil {
		c.String(http.StatusOK, "")
	}
	b := make([]byte, 0)
	b, _ = jsoniter.Marshal(b)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.String(http.StatusOK, string(b))
	return
}