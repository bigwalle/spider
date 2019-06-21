package http

import (
	bm "github.com/welcome112s/go-library/pkg/net/http/blademaster"
)

func testHello(c *bm.Context){
	c.JSON(helloSv.Hello(),nil)
}