package z

import (
	"fmt"
	"log"
	"net/http"
)

// 路由map的key格式
const routerMapKeyFormat = "[%s]%s"

// 处理过程间context传递
type ZContext struct {
	W    http.ResponseWriter
	R    *http.Request
	Data interface{}
}

// 子过程
type SubProcedure func(zCtx *ZContext) bool

// 过程
type Procedure []SubProcedure

// 路由
type Router struct {
	Pattern                 string    // 匹配路经
	Methods                 []string  // http方法集
	Procedure               Procedure // 过程集
	FirstChild, NextSibling *Router   // 子路由， 兄弟路由
}

// 创建路由
func NewRouter(pattern string, methods []string, procedure []SubProcedure) *Router {
	return &Router{
		Pattern:   pattern,
		Methods:   methods,
		Procedure: procedure,
	}
}

// 添加兄弟路由
func (r *Router) PushSibling(siblings ...*Router) {
	s := len(siblings)
	if s == 0 {
		return
	}
	r.NextSibling = siblings[0]
	if s > 1 {
		r.NextSibling.PushSibling(siblings[1:]...)
	}
}

// 添加子路由
func (r *Router) PushChild(children ...*Router) {
	c := len(children)
	if c == 0 {
		return
	}
	r.FirstChild = children[0]
	if c > 1 {
		r.FirstChild.PushSibling(children[1:]...)
	}
}

// 解析路由
func ParseRouter(r *Router, target *map[string]Procedure) {
	log.Println(r.Pattern, r.Methods)
	if r.FirstChild == nil {
		if len(r.Methods) == 0 {
			(*target)[r.Pattern] = r.Procedure
			return
		}
		for _, m := range r.Methods {
			key := fmt.Sprintf(routerMapKeyFormat, m, r.Pattern)
			(*target)[key] = r.Procedure
		}
		return
	}
	for c := r.FirstChild; c != nil; c = c.NextSibling {
		t := &Router{
			Pattern:    r.Pattern + c.Pattern,
			Methods:    filterMethods(r.Methods, c.Methods),
			Procedure:  append(r.Procedure, c.Procedure...),
			FirstChild: c.FirstChild,
		}
		ParseRouter(t, target)
	}
}

// 分层http方法过滤，存在于上层路由的http方法不存在于下层路由， 或存在于下层路由的http方法
// 但不存在于上层路由，都会在最终的路由结果中过滤。当前路由层中若其http方法集为空，表示本层
// 路由匹配由上层传导下来的所有http方法
func filterMethods(a, b []string) []string {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	t := make(map[string]struct{})
	r := make([]string, 0)
	for _, s := range a {
		t[s] = struct{}{}
	}
	for _, s := range b {
		if _, ok := t[s]; ok {
			r = append(r, s)
		}
	}
	return r
}

// 在传入路由上开启服务
func Run(r *Router, addr string) {
	rMap := make(map[string]Procedure)
	log.Println("parsing router...")
	ParseRouter(r, &rMap)
	log.Println("register handle func")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		method := request.Method
		mPath := fmt.Sprintf(routerMapKeyFormat, method, path)
		var (
			procedure Procedure
			ok        bool
		)
		if procedure, ok = rMap[path]; !ok {
			if procedure, ok = rMap[mPath]; !ok {
				writer.WriteHeader(400)
				return
			}
		}
		zCtx := &ZContext{
			W: writer,
			R: request,
		}
		for _, p := range procedure {
			if !p(zCtx) {
				return
			}
		}
	})
	log.Printf("starting server on %s...\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("starting server failed: %v\n", err)
	}
}
