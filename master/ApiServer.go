package master

import (
	"net"
	"net/http"
	"time"
)

var (
	G_apiServer *ApiServer
)

type ApiServer struct{
	httpServer *http.Server
}

func handleJobSave(w http.ResponseWriter, r *http.Request){

}

func InitApiServer() (err error){
	var (
		mux *http.ServeMux
		listener net.Listener
		httpServer *http.Server
	)

	//配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("/job/save",handleJobSave)

	//启动TCP监听
	if listener,err = net.Listen("tcp",":8070");err != nil{
		return
	}

	httpServer = &http.Server{
		ReadTimeout:5 * time.Second,
		WriteTimeout:5 * time.Second,
		Handler:mux,
	}

	G_apiServer = &ApiServer{
		httpServer: httpServer,
	}

	//启动服务端
	go httpServer.Serve(listener)

	return
}
