package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hxdyjx/Li/bootstrap"
	"github.com/hxdyjx/Li/servers/httpserver"
)

var appConfig = flag.String("conf", "conf/app.toml", "app config toml file")

func main() {
	// no show everyone debug info
	gin.SetMode("debug")
	//init conf
	flag.Parse()
	background := context.Background()
	bootstrap.MustLoad(*appConfig, background)

	// start http server
	err := httpserver.Run()
	if err != nil {
		panic(fmt.Sprintf("create httpserver error=%s\n", err.Error()))
	}

	// start socket server

}
