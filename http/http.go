package http

import (
	"spider/conf"
	"spider/service/test"
	bm "github.com/welcome112s/go-library/pkg/net/http/blademaster"
)

var helloSv *test.Service

func  Init(conf *conf.Config){
	initService(conf)
	engine:= bm.DefaultServer(conf.BM)
	initRouter(engine)
	if err := engine.Start(); err != nil {
		panic(err)
	}
	return

}

func initService(c *conf.Config){
	helloSv =test.New(c)
}


func initRouter(e *bm.Engine) {
	g := e.Group("/hello")
	{
		g.GET("/start", testHello)
	}
}