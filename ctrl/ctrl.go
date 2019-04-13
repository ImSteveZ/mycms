// ctrl包提供了http请求与响应的链式调用解决方案
// 注意，当前该包功能并发不安全，所以要避免并发的调用相关方法

package ctrl

import (
	"net/http"
	"sync"
)

// ctrl类型结构，实现http.Handler接口，封装当前handler需要执行的
// 过程集procedures，并包装了一个用于过程间进行通讯的上下文ctx
// 当类型的ServerHTTP方法被调用时, 将依次调用过程集procedures中的函数，直到
// 其中一个过程的返回值为false
type ctrl struct {
	sync.Mutex                                                              // 安全并发
	ctx        interface{}                                                  // 上下文
	procedures []func(interface{}, http.ResponseWriter, *http.Request) bool // 过程集
}

// Push方法为ctrl添加过程
func (ctrl *ctrl) Push(procedures ...func(interface{}, http.ResponseWriter, *http.Request) bool) *ctrl {
	ctrl.Lock()
	defer ctrl.Unlock()
	ctrl.procedures = append(ctrl.procedures, procedures...)
	return ctrl
}

// ServeHTTP方法实现http.Hander接口
func (ctrl *ctrl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctrl.Lock()
	defer ctrl.Unlock()
	for _, p := range ctrl.procedures {
		if !p(&ctrl.ctx, w, r) {
			return
		}
	}
	return
}

// Newctrl函数返回一个包装了传入ctx的ctrl类型指针
func NewCtrl(ctx interface{}) *ctrl {
	return &ctrl{ctx: ctx}
}
