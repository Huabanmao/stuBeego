package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "study.beego/models"
	_ "study.beego/routers"
)

func main() {
	beego.Run()
}
