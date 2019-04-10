package main

import (
	"log"
	"mycms/ctrls"
	"mycms/utils"
	"net/http"
)

func main() {
	// Public
	public := http.FileServer(http.Dir("./public"))
	http.Handle("/", http.StripPrefix("/", public))

	// Bind handlers to route
	http.HandleFunc("/usr/signUp", Middleware(ctrls.SignUpCtrl))
	http.HandleFunc("/usr/list", Middleware(ctrls.ListUserCtrl))

	// Listen http server
	log.Println("server start at 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// Middleware
func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Log
		utils.EnterLog(r)
		// Cross origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}