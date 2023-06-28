package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"study.beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.IndexController{}, "get:Get;post:Post")

	beego.Router("/mysql", &controllers.MysqlController{}, "get:MysqlGet")

	beego.Router("/insert", &controllers.ShowOrmController{}, "get:InsertOrm")
	beego.Router("/delete", &controllers.ShowOrmController{}, "get:DelOrm")
	beego.Router("/update", &controllers.ShowOrmController{}, "get:UpdateOrm")
	beego.Router("/query", &controllers.ShowOrmController{}, "get:QueryOrm")
}
