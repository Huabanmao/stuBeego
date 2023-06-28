package controllers

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"study.beego/models"
)

type ShowOrmController struct {
	beego.Controller
}

func (this *ShowOrmController) InsertOrm() {
	//1. 获取orm对象
	newOrm := orm.NewOrm()
	//2. 插入对象
	user := models.User{}
	user.Name = "小甜甜"
	//3. 执行插入操作
	insert, err := newOrm.Insert(&user)
	if err != nil {
		logs.Info("插入数据错误：", err)
		return
	}
	logs.Info("插入数据返回值：", insert)
	//4. 没输出，打印一个数据到浏览器
	this.Ctx.WriteString("插入数据成功")
}

func (this *ShowOrmController) DelOrm() {
	newOrm := orm.NewOrm()
	user := models.User{Id: 1}
	i, err := newOrm.Delete(&user)
	if err != nil {
		logs.Info("删除出错误：", err)
		return
	}
	logs.Info("删除返回值：", i)
	this.Ctx.WriteString("删除数据成功")
}

func (this *ShowOrmController) UpdateOrm() {
	ormer := orm.NewOrm()
	user := models.User{}
	user.Name = "牛夫人"
	user.Id = 5
	update, err := ormer.Update(&user)
	if err != nil {
		logs.Info("更新数据错误：", err)
		return
	}
	logs.Alert("更新数据返回值：", update)
	this.Ctx.WriteString("更新数据成功...")
}

func (this *ShowOrmController) QueryOrm() {
	ormer := orm.NewOrm()
	user := models.User{Name: "小甜甜"}
	err := ormer.Read(&user)
	if err != nil {
		logs.Info("查询表错误: ", err)
		return
	}
	logs.Info("查询的名字为：", user.Name)
	this.Ctx.WriteString("查询成功...")

}
