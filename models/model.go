package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
}

func init() {
	// 参数1   driverName
	// 参数2   数据库类型
	// 这个用来设置 driverName 对应的数据库类型
	// mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	//1. 连接数据库
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		logs.Info("数据库驱动错误: ", err)
		return
	}
	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30
	err = orm.RegisterDataBase("default", "mysql", "daxia:daxia@tcp(192.168.100.104:3306)/tanrui?charset=utf8", maxIdle, maxConn)
	if err != nil {
		logs.Info("数据库连接失败：", err)
		return
	}
	//2. 注册表
	orm.RegisterModel(new(User))
	//3. 生成表
	orm.RunSyncdb("default", false, true)
}
