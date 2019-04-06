package main

import (
	"log"
	"mycms/ctrls"
	"net/http"
)

func main() {
	// Bind handlers to route
	// http.HandleFunc("/", ctrls.IndexCtrl)
	http.HandleFunc("/usr/signup", ctrls.SignUpCtrl)
	http.HandleFunc("/usr/list", ctrls.ListUserCtrl)

	// Lisen http
	log.Println("server start...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
