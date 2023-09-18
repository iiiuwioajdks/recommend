package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	UserMysqlClient *sql.DB
)

func CreateMysqlClient() {
	createUserMysqlClient()
}

func createUserMysqlClient() {
	// 连接数据库
	var err error
	UserMysqlClient, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/user_profile")
	if err != nil {
		panic(err.Error())
	}

	// 测试连接是否成功
	err = UserMysqlClient.Ping()
	if err != nil {
		panic(err.Error())
	}
}
