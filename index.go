package main

import (
	"log"
	"mycms/ctrls"
	"mycms/utils"
	"mygo"
	"net/http"
)

func main() {
	myx := mygo.NewMyx()

	// Public
	myx.ServeFile("/", "./public")

	// Bind handlers to route
	myx.HandleFunc("/usr/signUp", EnterLoger, AllowOrigin, ctrls.SignUpCtrl)
	myx.HandleFunc("/usr/list", EnterLoger, AllowOrigin, ctrls.ListUserCtrl)

	// Listen http server
	log.Println("server start at 3000...")
	log.Fatal(http.ListenAndServe(":3000", myx))
}

func EnterLoger(c *mygo.Ctx, w http.ResponseWriter, r *http.Request) bool {
	utils.EnterLog(r)
	return true
}

func AllowOrigin(c *mygo.Ctx, w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return true
}
