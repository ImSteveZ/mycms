package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type DBCtxKey struct{}

const (
	Username     = "root"
	Password     = "hklmtt01#MYSQL"
	Protocol     = "tcp"
	Host         = "127.0.0.1"
	Port         = "3306"
	Database     = "mycms"
	MaxLifeTime  = 100
	MaxOpenConns = 16
	MaxIdleConns = 16
)

func NewMysqlDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		Username,
		Password,
		Protocol,
		Host,
		Port,
		Database,
	)
	mysqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("mysql init: %v", err))
	}
	mysqlDB.SetConnMaxLifetime(MaxLifeTime * time.Second)
	mysqlDB.SetMaxOpenConns(MaxOpenConns)
	mysqlDB.SetMaxIdleConns(MaxIdleConns)
	return mysqlDB, nil
}
