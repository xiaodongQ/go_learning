package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	// retCh := make(chan string) // 若用不缓冲的chan，则发送后时会阻塞到 读取这个chan的程序接收后 为止，即retCh <- ret，读取结束，执行下一条
	retCh := make(chan string, 1) // 若用缓冲chan，有足够缓冲时，顺序是 retCh <- ret，执行下一条指令(而不用管是否读取了channel)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

//
func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)        // 此处会阻塞到retCh通道能读取到数据为止
	time.Sleep(time.Second * 1) // 此处的sleep是要等待线程处理完
}

func TestAsynServiceWaitGroup(t *testing.T) {
	retCh := AsyncService()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		otherTask() //将调用调整为协程执行
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println(<-retCh) // 此处会阻塞到retCh通道能读取到数据为止，调整为协程
		wg.Done()
	}()

	wg.Wait()

	// 此处的sleep是要等待线程处理完，如果不加WaitGroup并注释该句，则协程执行otherTask时(其中做了sleep)，未完成就进程就结束了，没有等到打印"Task is done."
	// time.Sleep(time.Second * 1)
}
