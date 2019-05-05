package main

import (
	_ "sdrms/routers"
	"github.com/astaxie/beego"
	_ "sdrms/sysinit"
)

func main() {
	beego.Run("127.0.0.1:8080")

}

