package z

import (
	"context"
	"net/http"
)

type Middle func(ctrl Ctrl) Ctrl

type HttpMe struct {
	W http.ResponseWriter
	R *http.Request
}

type Ctrl struct {
	Pattern string
	Function func(ctx context.Context)
}
func Specify(prefix string, middles []Middle, ctrls[]Ctrl) {
	for _, ctrl := range ctrls {
		p := prefix + ctrl.Pattern
		handler := func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(context.Background(), "http", &HttpMe{w, r})
			for _, middle := range middles {
				ctrl = middle(ctrl)
			}
			ctrl.Function(ctx)
		}
		http.HandleFunc(p, handler)
	}
}