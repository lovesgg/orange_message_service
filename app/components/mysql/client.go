package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	config2 "orange_message_service/app/components/config"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetDb() *sql.DB {
	config := config2.GetConfig()
	userName := config.GetString("mysql.userName")
	password := config.GetString("mysql.password")
	ip := config.GetString("mysql.ip")
	dbName := config.GetString("mysql.dbName")

	//连接数据库
	db, err := sql.Open("mysql", userName+":"+password+"@tcp("+ip+")/"+dbName+"?charset=utf8")
	checkErr(err)
	return db
}
