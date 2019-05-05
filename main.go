package main

import (
	_ "ARTS-daka/routers"
	"github.com/astaxie/beego"
	_ "ARTS-daka/sysinit"
)

func main() {
	beego.Run("127.0.0.1:8080")
}

