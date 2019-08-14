package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	t.Run("prints 5 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{}) //结果写buffer中，但Countdown中调用的Sleep()方法则为L64定义的方法

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter) //这里有点不明包具体原理，但是执行过程应该是Fprinln实际上被L68的Write()方法替换，Sleep()同上一个子测试

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep} //初始化ConfigurableSleeper类型的结构体变量
	sleeper.Sleep()                                          //实际就是执行SpyTime类型的Sleep()方法

	if spyTime.durationSlept != sleepTime { //这个测试主要是看ConfigurableSleeper能不能正确配置时间长度
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep) //定义了结构体CountdownOperationsSpy的Sleep()方法，该方法执行的效果是向结构体的域，即一个字符串中写字符串常量sleep
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write) //定义了结构体CountdownOperationsSpy的Write()方法
	return
}

const write = "write"
const sleep = "sleep"

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) { //这个方法的定义是为了满足ConfigurableSleeper的第二个域对函数的格式要求
	s.durationSlept = duration //该方法实际就是修改结构体SpyTime的域的值
}
