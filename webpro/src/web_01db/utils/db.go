package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	err error
)

func init(){
	Db, err = sql.Open("mysql","root:123456@/test")
	Db.Exec()
	if err != nil{
		panic(err.Error())
	}
	fmt.Println(Db)
}

