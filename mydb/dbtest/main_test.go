package mydb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	mydb "github.com/RaymondCode/simple-demo/mydb/sqlc"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	dbSource = "root:baobao@tcp(81.68.118.43:3306)/dousheng"
)

var testQueries *mydb.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	testQueries = mydb.New(conn)
	fmt.Println("数据库连接成功")
	os.Exit(m.Run())
}
