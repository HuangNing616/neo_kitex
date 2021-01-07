package main

import (
	"fmt"
	"time"
)

/**
	Author: huangning
	Date: 2020/12/30
	Target: 了解channel
	Note:
	1. <-总是优先和最左边的类型结合
	2. channel关闭之后就无法向其发送数据
	3. Channel数据顺序
       Channel是一个FIFO(先进先出)队列，向channel发送数据的顺序和从channel接受数据的顺序一致
*/

func main() {

	// 1.初始化channel
	// 容量(capacity)代表Channel容纳最多的元素数量，代表Channel的缓存大小
	// 元素数量超过缓存大小，向channel发送数据会阻塞
	// 元素数量为空，从channel接收数据会阻塞
	// int表示chan里面的数据类型
	ch1 := make(chan int, 100)
	ch2 := make(chan int)
	ch3 := make(chan interface{}, 10)

	// 3.Channel closed
	// Channel关闭之后，继续向Channel发送数据会panic
	// 从Channel接受数据不会panic，会立即返回，返回的是chan数据类型的零值
	ch1 <- 12
	close(ch1)
	close(ch2)

	// ch1 <- 1 会报panic: send on closed channel
	a := <-ch1
	b,ok := <-ch2
	fmt.Printf("从ch1接收数据:%v\n", a)
	fmt.Printf("从ch2接收chan零值:%v, ch2是否关闭:%v\n", b, ok)

	// 4.Channel遍历
	// 向有缓冲区的channel发送数据, 被close之后依然会输出channel中的元素
	for i := 0; i < 10; i = i + 1 {
		ch3 <- i
	}
	close(ch3)
	for i := range ch3 {
		fmt.Printf("%v\n", i)
	}
	fmt.Println("=====根据FIFO, 结果按顺序输出======")

	// 向没有缓冲区的channel发送数据
	ch4 := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			ch4 <- i
		}
		// 必须close，否则遍历channel后会panic
		// 会提示：fatal error: all goroutines are asleep - deadlock!
		close(ch4)
	}()
	for i := range ch4 {
		fmt.Println(i)
	}
	fmt.Println("=====根据FIFO, 结果按顺序输出======")

	// 5. select语句
	cc := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("------", <-cc)
		}
		quit <- 0
	}()
	fibonacci(cc, quit)


	// 6. Channel用于timeout
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}
	fmt.Println("===============")

	// 7. goroutine同步
	// 利用 channel 我们可以实现goroutine之间相互通信
	done := make(chan bool, 1)
	go worker(done)

	// 等待任务完成
	fmt.Printf("结果:%v", <-done)

}

func worker(done chan bool) {

	time.Sleep(time.Second)
	done <- true
}

func fibonacci(c chan int, quit chan int) {
	x, y := 0, 1
	// 如果default都没有，则select语句会一直阻塞，直到至少有一个IO操作可以进行
	for {
		select {
		case c<- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}