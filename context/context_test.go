package context3

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	data := "123456789"

	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		// Server()函数返回一个可用作HTTP handler的HandlerFunc类型自定义普通函数
		svr := Server(store)

		// NewRequest方法返回一个新的可传入服务器的请求，可用于传入到一个http.Handler进行测试
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		// NewRecorder方法返回一个初始化的ResponseRecorder
		response := httptest.NewRecorder()

		// ServeHTTP方法的作用是进行函数调用：svr(response, request)。具体过程即：调用store.Fetch(r.Context())和执行Fprint写操作
		svr.ServeHTTP(response, request)

		// 判断response是否写入正确
		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		/*
		 * WithCancel方法接受一个Context并返回一个Context（即第一个返回值），返回的Context里面包含一个Done() channel
		 * 当调用返回的cancel函数（即第二个返回值）或父Context的Done channel被关闭时，返回的Context（第一个返回值）的Done channel被关闭
		 */
		cancellingCtx, cancel := context.WithCancel(request.Context())
		// AfterFunc方法会等待其第一个参数定义的时间流逝后，调用第二个参数传入的函数，此时即调用前面所述的cancel函数
		time.AfterFunc(T1, cancel)
		// WithContext方法会返回调用者（即等号右侧request）的一个浅拷贝，该拷贝中包含了实参发生的context修改
		request = request.WithContext(cancellingCtx)

		// httptest.NewRecorder的返回值是 *ResponseRecorder， 同理自定义的SpyResponseWriter赋给response的初始化值也应是指针类型
		response := &SpyResponseWriter{}

		/*
		 * 这里就是通过自定义的SpyResponseWriter实现响应结果为将written域写为true
		 * 测试的期望结果是cancel执行成功，即不会写任何的response，因此written期望值应为false
		 */
		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})
}
