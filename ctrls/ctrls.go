package ctrls

import (
	"fmt"
	"io"
	"mycms/utils"
	"mycms/modls"
	"net/http"
	"strings"
)

// SignUpReq is sign up request data struct
type SignUpReq struct {
	UserName, Email, Password, RepeatPassword string
}

// SginUpResp struct
type SignUpResp struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data SignUpData `json:"data"`
}

// SignUpValidateData
type SignUpData struct {
	UserID int64 `json:"user_id,omitempty"`
	ValidateData ValidateData `json:"validate_data,omitempty"`
}

// ValidateData
type ValidateData struct {
	UserName FieldValidate `json:"user_name,omitempty"`
	Email FieldValidate `json:"email,omitempty"`
	Password FieldValidate`json:"password,omitempty"`
	RepeatPassword FieldValidate `json:"repeat_password,omitempty"` 
}

// FieldValidate
type FieldValidate struct {
	Value, Tip string
}

// SignUpCtrl
func SignUpCtrl(w http.ResponseWriter, r *http.Request) {
	utils.EnterLog(r)
	// Invalid request method
	if r.Method != "POST" {
		result := &SignUpResp{
			Status: 400,
			Message: fmt.Sprintf("invalid request method: %s", r.Method),
		}
		utils.ServeJson(w, result)
		return
	}
	// Extract signUpData
	var signUpData SignUpReq
	if err := utils.ExtractJson(r, &signUpData); err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Data validate
	var result = &SignUpResp{}
	if validateData, ok := validateSignUpData(signUpData); !ok {
		result.Status = 300
		result.Message = "form data is invalid"
		result.Data.ValidateData = validateData
		utils.ServeJson(w, result);
		return
	}

	// User
	password, passwordSalt := utils.Password(signUpData.Password)
	user := &modls.User{
		Email: signUpData.Email,
		UserName: signUpData.UserName,
		Password: password,
		PasswordSalt: passwordSalt,
	}
	// New user model
	userModl, err := modls.NewModl()
	if err != nil {
		result.Status = 400
		result.Message = "database connect error"
		utils.ServeJson(w, result)
		return
	}
	// Add user
	userID, err := userModl.AddOrUpdateUser(user)
	if err != nil || userID == 0 {
		result.Status = 400
		result.Message = "database operation error"
		utils.ServeJson(w, result)
		return
	}
	// Success
	result.Status = 200
	result.Message = "register success"
	result.Data.UserID = userID
	utils.ServeJson(w, result)
	return
}

func validateSignUpData(signUpData SignUpReq) (validateData ValidateData, ok bool) {
	ok = true
	// UserName
	if strlen := len(signUpData.UserName); strlen < 2 || strlen > 10 {
		validateData.UserName = FieldValidate{
			Value: signUpData.UserName,
			Tip: "username's lenth can not less than 2 or more than 10",
		}
		ok = false
	}
	// Email
	if !strings.Contains(signUpData.Email, "@") {
		validateData.Email = FieldValidate{
			Value: signUpData.Email,
			Tip: "email's format is invalid, e.g.demo@mycms.com",
		}
		ok = false
	}
	// Password
	if strlen := len(signUpData.Password); strlen < 6 || strlen > 24 {
		validateData.Password = FieldValidate{
			Value: signUpData.Password,
			Tip: "password's lenth can not less than 6 or more than 24",
		}
		ok = false
	}
	// RepeatPassword
	if signUpData.Password != signUpData.RepeatPassword {
		validateData.RepeatPassword = FieldValidate{
			Value: signUpData.RepeatPassword,
			Tip: "two twice password are not equal",
		}
		ok = false
	}
	return
}