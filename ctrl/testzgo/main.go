package main

import (
	"mycms/ctrl"
	"net/http"
	"encoding/json"
)

type Ctx struct {
	Message string
}

func (ctx *Ctx) ServeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	var ctx Ctx
	http.Handle("/", ctrl.NewCtrl(ctx).Push(auth, index))
	http.Handle("/usr", ctrl.NewCtrl(ctx).Push(auth, usr))
	http.ListenAndServe(":3000", nil)
}

func auth(ctx interface{}, w http.ResponseWriter, r *http.Request) bool {
	cctx := ctx.(*Ctx)
	cctx.Message += "auth "
	return true
}

func index(ctx interface{}, w http.ResponseWriter, r *http.Request) bool {
	cctx := ctx.(*Ctx)
	cctx.Message += "index"
	w.Write([]byte(cctx.Message))
	return true
}

func usr(ctx interface{}, w http.ResponseWriter, r *http.Request) bool {
	cctx := ctx.(*Ctx)
	cctx.Message += "usr"
	w.Write([]byte(cctx.Message))
	return true
}