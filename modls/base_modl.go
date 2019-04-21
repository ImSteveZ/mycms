package modls

import (
	"database/sql"
	"db"
	_ "github.com/go-sql-driver/mysql"
)

type BaseModl struct {
	MysqlDB *sql.DB
}

var UserModl UserModl

func init() {
	UserModl = &UserModl{
		MysqlDB: db.MysqlDB
	}
}
