package "db"

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

const (
	USERNAME = "root"
	PASSWORD = "hklmtt01#MYSQL"
	PROTOCOL = "tcp"
	HOST = "127.0.0.1"
	PORT = "3306"
	DATABASE = "mycms"
)

func init() {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s",
		USERNAME, PASSWORD, PROTOCOL, HOST, PORT, DATABASE,
	)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("database init fail: %v\n", err)
		return
	}
	DB.SetConnMaxLifetime(100 * time.Second)
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	log.Println("database init success...")
}
