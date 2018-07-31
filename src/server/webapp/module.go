package webapp

import (
	"encoding/json"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/module/base"
)

type WebApp struct {
	basemodule.BaseModule
}

var Module = func() module.Module {
	webapp := new(WebApp)
	return webapp
}

func (m *WebApp) GetType() string {
	//很关键,需要与配置文件中的Module配置对应
	return "Webapp"
}
func (m *WebApp) Version() string {
	//可以在监控时了解代码版本
	return "1.0.0"
}
func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		//[26/Oct/2017:19:07:04 +0800]`-`"GET /g/c HTTP/1.1"`"curl/7.51.0"`502`[127.0.0.1]`-`"-"`0.006`166`-`-`127.0.0.1:8030`-`0.000`xd
		log.Info("%s %s %s [%s] in %v", r.Method, r.URL.Path, r.Proto, r.RemoteAddr, time.Since(start))
	})
}
func (self *WebApp) OnInit(app module.App, settings *conf.ModuleSettings) {
	self.BaseModule.OnInit(self, app, settings)
}

func Statushandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
	})
}
func (m *WebApp) Run(closeSig chan bool) {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Error("webapp server error", err.Error())
		return
	}
	go func() {
		log.Info("webapp server Listen : %s", ":8080")
		root := mux.NewRouter()
		status := root.PathPrefix("/status")
		status.HandlerFunc(Statushandler)

		static := root.PathPrefix("/mqant/")
		static.Handler(http.StripPrefix("/mqant/", http.FileServer(http.Dir(m.GetModuleSettings().Settings["StaticPath"].(string)))))
		ServerMux := http.NewServeMux()
		ServerMux.Handle("/", root)
		http.Serve(listen, loggingHandler(ServerMux))
	}()

	<-closeSig
	log.Info("webapp server Shutting down")
	listen.Close()
}
func (self *WebApp) OnDestroy() {
	//一定别忘了关闭RPC
	self.GetServer().OnDestroy()
}
