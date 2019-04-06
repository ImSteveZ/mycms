package ctrls

import (
	"fmt"
	"log"
	"mycms/db"
	"mycms/modls"
	"mycms/utils"
	"net/http"
	"strings"
)

// Init single modl
var userModl *modls.Modl

func init() { userModl = modls.NewModl(db.DB) }

// SignUpReq is sign up request data struct
type SignUpReq struct {
	UserName, Email, Password, RepeatPassword string
}

// Resp struct
type Resp struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SignUpValidateData
type SignUpData struct {
	UserID        int64          `json:"user_id,omitempty"`
	ValidateInfos []ValidateInfo `json:"validate_infos,omitempty"`
}

// FieldValidate
type ValidateInfo struct {
	Field, Value, Tip string
}

// ListUserCtrl
func ListUserCtrl(w http.ResponseWriter, r *http.Request) {
	var resp *Resp
	users, err := userModl.ListUsers()
	// Failed
	if err != nil {
		log.Printf("list user: %v\n", err)
		resp = &Resp{
			Status:  400,
			Message: "query failed",
		}
		utils.ServeJson(w, resp)
		return
	}
	// Success
	resp = &Resp{
		Status:  200,
		Message: "query success",
		Data:    users,
	}
	utils.ServeJson(w, resp)
	return
}

// SignUpCtrl
func SignUpCtrl(w http.ResponseWriter, r *http.Request) {
	utils.EnterLog(r)
	var resp *Resp

	// Invalid request method
	if r.Method != "POST" {
		resp = &Resp{
			Status:  400,
			Message: fmt.Sprintf("invalid request method: %s", r.Method),
		}
		utils.ServeJson(w, resp)
		return
	}

	// Extract request data
	var signUpData SignUpReq
	utils.ExtractJson(r, &signUpData)

	// Validate request data
	if validateInfos, ok := validateSignUpData(signUpData); !ok {
		resp = &Resp{
			Status:  300,
			Message: "form data is invalid",
			Data: SignUpData{
				ValidateInfos: validateInfos,
			},
		}
		utils.ServeJson(w, resp)
		return
	}

	// User
	password, passwordSalt := utils.Password(signUpData.Password)
	user := &modls.User{
		Email:        signUpData.Email,
		UserName:     signUpData.UserName,
		Password:     password,
		PasswordSalt: passwordSalt,
	}

	// Add user
	userID, err := userModl.AddOrUpdateUser(user)
	if err != nil || userID == 0 {
		log.Printf("add user failed: %v\n", err)
		resp = &Resp{
			Status:  400,
			Message: "system operation error",
		}
		utils.ServeJson(w, resp)
		return
	}

	// Register success
	resp = &Resp{
		Status:  200,
		Message: "register success",
		Data: SignUpData{
			UserID: userID,
		},
	}
	utils.ServeJson(w, resp)
	return
}

func validateSignUpData(signUpData SignUpReq) (validateInfos []ValidateInfo, ok bool) {
	ok = true
	// UserName
	if strLen := len(signUpData.UserName); strLen < 2 || strLen > 10 {
		validateInfos = append(validateInfos, ValidateInfo{
			"user_name",
			signUpData.UserName,
			"username's lenth must between 2 and 10",
		})
		ok = false
	}
	// Email
	if !strings.Contains(signUpData.Email, "@") {
		validateInfos = append(validateInfos, ValidateInfo{
			"email",
			signUpData.Email,
			"email's format must like 'demo@mycms.com'",
		})
		ok = false
	}
	// Password
	if strlen := len(signUpData.Password); strlen < 6 || strlen > 24 {
		validateInfos = append(validateInfos, ValidateInfo{
			"password",
			signUpData.Password,
			"password's lenth must between 6 and 24",
		})
		ok = false
	}
	// RepeatPassword
	if signUpData.Password != signUpData.RepeatPassword {
		validateInfos = append(validateInfos, ValidateInfo{
			"repeat_password",
			signUpData.RepeatPassword,
			"password and repeat_password must be equal",
		})
		ok = false
	}
	return
}
