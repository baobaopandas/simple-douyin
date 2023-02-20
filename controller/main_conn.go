package controller

import (
	"database/sql"
	"fmt"

	"github.com/RaymondCode/simple-demo/config"
	mydb "github.com/RaymondCode/simple-demo/mydb/sqlc"
	_ "github.com/go-sql-driver/mysql"
)

// dbSource := "root:baobao@tcp(81.68.118.43:3306)/dousheng"

func GetConn() *mydb.Queries {
	dbConfig := config.CONFIG.DbConfig
	dbDriver := dbConfig.DbDriver
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		fmt.Println("cannot connect to db: ", err)
		return nil
	} else {
		fmt.Println("链接成功")
	}
	Queries := mydb.New(conn)
	return Queries
}
