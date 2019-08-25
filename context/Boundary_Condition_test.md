以下为程序执行关键路径上的部分代码：

* `contxet_text.go` 中调用cancel的部分

```go
...
cancellingCtx, cancel := context.WithCancel(request.Context())
// Cancel调用延迟
time.AfterFunc(T1, cancel)
...
svr.ServeHTTP(response, request)
...
```

* `testdouble.go` 中Fetch方法的实现

```go
func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
  data := make(chan string, 1)

  //go func执行延迟
	time.Sleep(T2)
	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store got cancelled")
        // extra log before return
        fmt.Println("cancelling")
				return
			default:
        //default执行延迟
				time.Sleep(T3)
				result += string(c)
        // extra log here
        fmt.Println(string(c))
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
    // extra log before return
    fmt.Println("Done")
		return "", ctx.Err()
	case res := <-data:
    // extra log before return
    fmt.Println("Response")
		return res, nil
	}
}
```

* Fetch方法在执行语句`svr.ServeHTTP()` 时被调用，Fetch方法内部会创建一个goroutine，通过select监听来决定执行Cancellation还是ResponseWrite。



> time.Sleep(N * time.Millisecond)

通过增加延迟来控制程序块的执行顺序，下表为不同配置下的测试结果。

| 测试编号 | cancel延迟 T1 | go func延迟 T2 | default延迟 T3 | 测试结果 |
| :------: | :-----------: | :------------: | :------------: | :------: |
| original |       5       |       0        |       10       |          |
|    1     |       0       |       0        |       0        |   随机   |
|    2     |       0       |       0        |       1        |          |
|    3     |      10       |       0        |       3        |          |
|    4     |       5       |       8        |       10       |          |



大致流程：

* go func执行后，select会监听 `ctx.Done()` ，如果没有close则执行default后的代码块。
* default代码块中会先执行一个延时，以模拟slow process。如果在延迟结束前cancel函数还未调用，则继续执行default代码块后续的部分；如果在default的延迟结束前执行了cancel函数，则执行cancellation。