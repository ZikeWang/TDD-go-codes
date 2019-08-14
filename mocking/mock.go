package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper allows you to put delays
type Sleeper interface {
	//定义接口类型，后面Countdown函数中第二个形参类型为Sleeper，则定义了Sleep()方法的类型都可作为实参传入
	//如L17定义类型后，在L24定义Sleep()方法，然后在L45作为实参传入
	//又如Test中：（L60，L64，L14/L29）
	Sleep()
}

// ConfigurableSleeper is an implementation of Sleeper with a defined delay
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep will pause execution for the defined Duration
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration) //为结构体类型ConfigurableSleeper定义Sleep方法：将第一个域作为参数来触发第二个域
}

const finalWord = "Go!"
const countdownStart = 3

// Countdown prints a countdown from 5 to out with a delay between count provided by Sleeper
func Countdown(out io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep} //这里是使用标准库中的Sleep函数和时间参数
	Countdown(os.Stdout, sleeper)                                //使用标准输出
}
