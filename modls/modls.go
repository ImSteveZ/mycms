package modls

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = "root"
	PROTOCOl = "tcp"
	HOST     = "localhost"
	PORT     = "3306"
	DATABASE = "mycms"
)

type modl struct {
	*sql.DB
}

func NewModl() (*modl, error) {
	dsn = fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s",
		USERNAME, PASSWORD, PROTOCOL, HOST, PROT, DATABASE,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifeTime(100 * time.second) // 最大连接周期，超时即close
	db.SetMaxOpenConns(100)                  // 最大连接数
	db.SetMaxIdleConns(16)                   // 最大闲置连接数
	return &modl{DB: db}
}

func (modl *modl) GetUserNameCount(username string) (bool, error) {
	var querysql := `select count(1) from users where UserName = ?`
	stmt, err := modl.Prepare(querysql)
	defer stmt.Close()
	if err != nil {
		return false, err
	}
	rows, err := stmt.Query(username)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	for rows.Next() {
		// TODO
	}
	// TODO
}











