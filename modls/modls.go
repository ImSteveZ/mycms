package modls

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	USERNAME = "root"
	PASSWORD = "htlmtt01#MYSQL"
	PROTOCOl = "tcp"
	HOST     = "localhost"
	PORT     = "3306"
	DATABASE = "mycms"
)

type modl struct {
	DB *sql.DB
}

type User struct {
	ID int64
	Email string
	UserName string
	Password string
	PasswordSalt string
}
	
func NewModl() (*modl, error) {
	dsn := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s",
		USERNAME, PASSWORD, PROTOCOl, HOST, PORT, DATABASE,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(100 * time.Second) // 最大连接周期，超时即close
	db.SetMaxOpenConns(100)                  // 最大连接数
	db.SetMaxIdleConns(16)                   // 最大闲置连接数
	return &modl{DB: db}, nil
}

// AddOrUpdateUser
func (modl *modl) AddOrUpdateUser(user *User) (id int64, err error) {
	addSql := `insert into users (UserName, Email, Password, PasswordSalt, CreatedOn)
		values (?, ?, ?, ?, ?)
		where not exists (
			select 1 from users where Email = ?
			and IsDeleted = 0
		)
	`
	var stmt *sql.Stmt
	stmt, err = modl.DB.Prepare(addSql)
	if err != nil {
		return
	}
	defer stmt.Close()
	var rst sql.Result
	createOn := time.Now().Format("2006-01-02 15:04:05")
	rst, err = stmt.Exec(user.UserName,
		user.Email,
		user.Password,
		user.PasswordSalt,
		createOn,
	)
	if err != nil {
		return
	}
	return rst.LastInsertId()
}

// GetEmailCount 
//func (modl *modl) GetEmailCount(email string) (count int, err error) {
//	querySql := `select count(1) from users where Email = ?
//		and IsDeleted = 0`
//	var stmt *sql.Stmt
//	stmt, err = modl.DB.Prepare(querySql)
//	if err != nil {
//		return
//	}
//	defer stmt.Close()
//
//	var rst sql.Result
//	rst, err = stmt.Exec(email)
//	if err != nil {
//		return
//	}
//
//	err = rst.Scan(&count)
//	return
//}