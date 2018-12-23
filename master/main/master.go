package main

import (
	"fmt"
	"github.com/distributed/crontab/master"
	"runtime"
)

func main(){
	var (
		err error
	)

	//初始化线程
	initEnv()

	//启动Api HTTP服务
	if err = master.InitApiServer(); err != nil{
		goto ERR
	}

	return

	ERR:
		fmt.Println(err)
}

func initEnv(){
	runtime.GOMAXPROCS(runtime.NumCPU())
}
