package httpserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hxdyjx/Li/library/resource"
)

func SetRouter() *gin.Engine {
	var engine *gin.Engine
	if gin.Mode() == "debug" {

	} else {

	}
	engine = gin.Default()
	// insert middleware
	//engine.Use()
	return engine
}

func Run() error {
	engine := SetRouter()
	SetRemoteRouter(engine)
	err := engine.Run(fmt.Sprintf("%s:%d", resource.AppConf.Web.Address, resource.AppConf.Web.Port))
	if err != nil {
		return err
	}
	return nil
}
