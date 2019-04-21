package modls

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// Modl
type Modl struct{ DB *sql.DB }

// NewModl
func NewModl(db *sql.DB) *Modl { return &Modl{DB: db} }

// User
type User struct {
	ID           int64  `json:"user_id"`
	Email        string `json:"email"`
	UserName     string `json:"user_name"`
	Password     string `json:"password,omitempty"`
	PasswordSalt string `json:"password_salt,omitempty"`
}

// ListUser
func (modl *Modl) ListUsers() ([]*User, error) {
	listSql := `select ID, UserName, Email from users where 1=1`
	rows, err := modl.DB.Query(listSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*User
	for rows.Next() {
		var (
			id       int64
			userName string
			email    string
		)
		err := rows.Scan(&id, &userName, &email)
		if err != nil {
			return nil, err
		}
		users = append(users, &User{ID: id, UserName: userName, Email: email})
	}
	return users, nil
}

// AddOrUpdateUser
func (modl *Modl) AddOrUpdateUser(user *User) (id int64, err error) {
	addSql := `insert into users (UserName, Email, Password, PasswordSalt, IsApproved, IsLocked, CreatedOn)
		select ?, ?, ?, ?, ?, ?, ? from dual
		where not exists (
			select * from users where Email = ? 
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
	rst, err = stmt.Exec(
		user.UserName,
		user.Email,
		user.Password,
		user.PasswordSalt,
		0,
		1,
		createOn,
		user.Email,
	)
	if err != nil {
		return
	}
	return rst.LastInsertId()
}

// GetEmailCount
//func (modl *Modl) GetEmailCount(email string) (count int, err error) {
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
