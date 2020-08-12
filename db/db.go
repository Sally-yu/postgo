package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var DBE *xorm.Engine

const (
	host     = "112.126.57.214:5432"
	user     = "postgres"
	password = "aaaaaa"
	dbname   = "mydb"
)

func ConnectDB()  {
	var connectStr string
	connectStr="postgres://"+user+":"+password+"@"+host+"/"+dbname+"?sslmode=disable;"
	var err error
	DBE, err = xorm.NewEngine("postgres", connectStr)
	if err!=nil {
		fmt.Println(err.Error())
	}
}


