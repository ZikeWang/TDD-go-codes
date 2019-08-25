package context3

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
)

// SpyStore allows you to simulate a store and see how its used
type SpyStore struct {
	response string
	t        *testing.T
}

// Fetch returns response after a short delay
func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	time.Sleep(T2)
	go func() {
		var result string
		for _, c := range s.response {

			/*
			 * select中的case语句必须是一个channel操作
			 * 1. 如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。
			 * 2. 如果没有可运行的case语句，且有default语句，那么就会执行default的动作。
			 * 3. 如果没有可运行的case语句，且没有default语句，select将阻塞，直到某个case通信可以运行
			 */

			select {

			/*
			 * Done在工作结束时返回一个关闭的channel，这代表这个context应该被取消
			 * 如果context永远不会被取消则Done返回nil
			 * 当cancel函数被调用时，WithCancel会安排Done关闭
			 */

			case <-ctx.Done():
				s.t.Log("spy store got cancelled") //只有加上-v标志才会显示以下Log的打印内容
				fmt.Println("cancelling")
				return
			default:
				time.Sleep(T3)
				result += string(c)
				fmt.Println(string(c))
			}
		}
		data <- result
	}()

	select {
	//对Done的连续调用会返回相同的值
	case <-ctx.Done():
		fmt.Println("Done") //这里很奇怪的一点是加上8ms延迟后，（5ms，10ms）Done每次都在cancelling前面打印？？？？？？？？
		return "", ctx.Err()
	case res := <-data:
		fmt.Println("Response")
		return res, nil
	}
}

/*
 * 以下代码自定义了SpyResponseWriter结构体，并定义了满足 ResponseWriter 接口类型的三个方法。 ResponseWriter类型的定义如下：
 * type ResponseWriter interface {
 *     Header() Header
 *     Write([]byte) (int, error)
 *     WriteHeader(statusCode int)
 * }
 * 自定义的类型中所有方法都是结构体的written标志写为true，用于模拟完成了response的写入操作
 * 定义这样一个结构体及其方法的目的就是测试在cancel的情况下不会完成写入response的操作
 * 而httptest.ResponseRecorder不支持这样的测试，因为在v2.0版本中只是通过显式的在SpyStore结构体中将一个bool域cancel写true或者false来判断是否取消，无法判断ResponseWriter到底在cancel后是否有写过数据
 */

// SpyResponseWriter checks whether a response has been written
type SpyResponseWriter struct {
	written bool
}

// Header will mark written to true
func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

// Write will mark written to true
func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

// WriteHeader will mark written to true
func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}
