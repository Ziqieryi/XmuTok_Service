package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Link() {
	// 配置 MySQL 连接参数
	username := "root"        //账号
	password := "939791250wQ" //密码
	host := "127.0.0.1"       //数据库地址，可以是Ip或者域名
	port := 3307              //数据库端口
	Dbname := "douyin"        //数据库名

	// 通过前面的数据库参数，拼接 Mysql DSN，其实就是数据库连接串（数据源名称）
	// Mysql dsn格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
	// 类似{username}使用花括号包着的名字都是需要替换的参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)

	// 连接 Mysql
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
}
