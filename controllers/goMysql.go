package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlController struct {
	beego.Controller
}

func (m *MysqlController) MysqlGet() {
	//conn, _ := sql.Open("mysql", "daxia:daxia@tcp(192.168.100.104:3306)/tanrui?charset=utf8")
	//err := conn.Ping()
	//if err != nil {
	//	logs.Info("数据库连接错误：", err)
	//	return
	//}
	//defer conn.Close()
	////操作数据库
	//_, err = conn.Exec("create table if not exists stutest(id int auto_increment primary key, name varchar(10) not null, phone char(13))charset=utf8")
	//if err != nil {
	//	logs.Info("创建表失败：", err)
	//	m.Ctx.WriteString("创建数据库失败")
	//	return
	//}
	//m.Ctx.WriteString("创建数据库成功")

}
