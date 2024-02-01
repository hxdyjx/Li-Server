package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/hxdyjx/Li/servers/httpserver/server"
)

func SetRemoteRouter(e *gin.Engine) {
	v1 := e.Group("api/v1")
	v1.POST("/login", server.Login)
}
