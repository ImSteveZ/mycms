package main

import (
	"db"
	"log"
	"mycms/ctrls"
	"mycms/utils"
	"net/http"
	"nicego"
)

// CtxKey
type CtxKey string

func main() {
	sqlDB := db.NewDB()
	defer sqlDB.Close()

	// Route
	Ctx := context.WithValue(context.Background(), CtxKey("DB"), defaultDB)
	rt := nicego.NewRoute(ctx)

	// Register router
	rt.From("/").Use(Logger).Static("./public") // static files
	rt.From("/usr/signUP").Use(Logger, AllowOrigin).Do(ctrls.SignUpCtrl)
	rt.From("/usr/list").Use(Logger, AllowOrigin).Do(ctrls.ListUserCtrl)

	// Serve http
	log.Fatal(http.ListenAndServe(":3000", rt))
}

func EnterLoger(c *mygo.Ctx, w http.ResponseWriter, r *http.Request) bool {
	utils.EnterLog(r)
	return true
}

func AllowOrigin(c *mygo.Ctx, w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return true
}
