package mgate

import (
	"fmt"

	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/gate/base"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
)

var Module = func() module.Module {
	gate := new(Gate)
	return gate
}

type Gate struct {
	basegate.Gate
}

func (m *Gate) GetType() string {
	//很关键,需要与配置文件中的Module配置对应
	return "Gate"
}
func (m *Gate) Version() string {
	//可以在监控时了解代码版本
	return "1.0.0"
}

//当连接建立  并且MQTT协议握手成功
func (gate *Gate) Connect(session gate.Session) {
	log.Info("客户端建立了链接")
}

//当连接关闭	或者客户端主动发送MQTT DisConnect命令 ,这个函数中Session无法再继续后续的设置操作，只能读取部分配置内容了
func (gate *Gate) DisConnect(session gate.Session) {
	log.Info("客户端断开了链接")
}

/**
是否需要对本次客户端请求进行跟踪
*/
func (gate *Gate) OnRequestTracing(session gate.Session, topic string, msg []byte) bool {
	if session.GetUserid() == "" {
		//没有登陆的用户不跟踪
		return false
	}
	//if session.GetUserid()!="liangdas"{
	//	//userId 不等于liangdas 的请求不跟踪
	//	return false
	//}
	return true
}

func (gate *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
	gate.Gate.OnInit(gate, app, settings)
	gate.Gate.SetSessionLearner(gate)
	gate.Gate.SetStorageHandler(gate)
	gate.Gate.SetTracingHandler(gate)
}

func (gate *Gate) Storage(Userid string, session gate.Session) (err error) {
	log.Info("需要处理对Session的持久化")
	return nil
}

func (gate *Gate) Delete(Userid string) (err error) {
	log.Info("需要删除Session持久化数据")
	return nil
}

func (gate *Gate) Query(Userid string) (data []byte, err error) {
	log.Info("查询Session持久化数据")
	return nil, fmt.Errorf("no redis")
}

func (gate *Gate) Heartbeat(Userid string) {
	log.Info("用户在线的心跳包")
}
