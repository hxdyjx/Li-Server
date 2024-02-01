package bootstrap

import (
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/hxdyjx/Li/library/conf"
	"github.com/hxdyjx/Li/library/resource"
	"log/slog"
	"os"
	"time"
)

func MustLoad(confPath string, ctx context.Context) {
	MustInitAppConf(confPath)
	MustInitLog()
}

// MustInitAppConf prase app.toml
func MustInitAppConf(confPath string) {
	_, err := os.Stat(confPath)
	if err != nil {
		panic(fmt.Sprintf("don't found app.toml path:%v\n", err.Error()))
	}
	var app conf.App

	_, err = toml.DecodeFile(confPath, &app)
	if err != nil {
		panic(fmt.Sprintf("prase app.toml error=%v\n", err.Error()))
	}
	resource.AppConf = &app
}

// MustInitLog init log handle
func MustInitLog() {
	format := time.Now().Format("20060102")
	fHttp, err := os.Create("log/http/" + format + ".log")
	if err != nil {
		panic(fmt.Sprintf("create http log error=%s\n", err.Error()))
	}
	fSocket, err := os.Create("log/socket/" + format + ".log")
	if err != nil {
		panic(fmt.Sprintf("create socket log error=%s\n", err.Error()))
	}
	defer func() {
		fHttp.Close()
		fSocket.Close()
	}()

	resource.HttpLog = slog.New(slog.NewJSONHandler(fHttp, nil))
	resource.SocketLog = slog.New(slog.NewJSONHandler(fSocket, nil))

	resource.HttpLog.Info("create log successful")
	resource.SocketLog.Info("create log successful")
}
