package main

import (
	"algorithmplatform/bootstrap"
	"algorithmplatform/global"
)

func main() {

	// 初始化配置
	bootstrap.InitializeConfig()
	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()
	// 启动算法执行器
	bootstrap.InitializeActuator()
	// 启动服务器
	bootstrap.RunServer()
}
