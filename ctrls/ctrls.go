package ctrls

import (
	"fmt"
	"io"
	"log"
	"mycms/utils"
	"net/http"
	"strings"
)

// JsonResult is response data struct
type JsonResult struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SignUpReq is sign up request data struct
type SignUpReq struct {
	Username, Email, Password string
}

// Error format
const (
	ServeJsonErr   = "%s write json failed: %v\n"
	ExtractJsonErr = "%s extract json failed: %v\n"
)

// IndexCtrl
func IndexCtrl(w http.ResponseWriter, r *http.Request) {
	utils.EnterLog(r)
	w.Header().Set("Content-Type", "application/json")
	result := &JsonResult{Status: 200, Message: "Welcome to MyCMS"}
	if err := utils.ServeJson(w, result); err != nil {
		log.Printf(ServeJsonErr, "indexCtrl", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

// SignUpCtrl
func SignUpCtrl(w http.ResponseWriter, r *http.Request) {
	utils.EnterLog(r)
	w.Header().Set("Content-Type", "application/json")
	// Invalid request method
	if r.Method != "POST" {
		result := &JsonResult{
			Status: 400,
			Message: fmt.Sprintf("invalid request method: %s, POST is must needed", r.Method),
		}
		if err := utils.ServeJson(w, result); err != nil {
			log.Printf(ServeJsonErr, "signUpCtrl", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	// POST request
	var signUpData SignUpReq
	if err := utils.ExtractJson(r, &signUpData); (err != nil && err != io.EOF) {
		log.Printf(ExtractJsonErr, "signUpCtrl", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Data validate
	result := &JsonResult{}
	switch {
	// username
	case len(signUpData.Username) < 2:
		result.Message = "username is too short, at least 2"
	// password
	case len(signUpData.Password) < 6:
		result.Message = "password is too short, at least 6"
	// email
	case !strings.Contains(signUpData.Email, "@"):
		result.Message = "email is invalid, e.g.demo@mycms.com"
	}
	if result.Message != "" {
		result.Status = 400
	} else {
		result = &JsonResult{
			Status: 200,
			Message: "Success",
		}
	}
	if err := utils.ServeJson(w, result); err != nil {
		log.Printf(ServeJsonErr, "SginUpCtrl", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
