package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

func DoCommand() {
	ctx, cancel := context.WithCancel(context.Background())
	a :="asca"
	b :="asca"
	fmt.Println(a,b)
	go func(cancelFunc context.CancelFunc) {
		time.Sleep(3 * time.Second)
		cancelFunc()
	}(cancel)
	//err := Command(ctx, "ping www.baidu.com")
	//err := Command(ctx, "gospider -s \"https://mjt.emakerzone.com/origin/musiclist\" -o output -c 10 -d 1")
	err := Command(ctx, "gospider -q -s \"https://www.baidu.com\"")
	fmt.Println(err)
}

func Command(ctx context.Context, cmd string) error {
	sysType := runtime.GOOS
	c := &exec.Cmd{}
	if sysType == "windows" {
		c = exec.CommandContext(ctx, "cmd", "/C", cmd)
	}
	if sysType == "linux" || sysType == "mac" {
		c = exec.CommandContext(ctx, "bash", "-c", cmd)
	}
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			// 其实这段去掉程序也会正常运行，只是我们就不知道到底什么时候Command被停止了，而且如果我们需要实时给web端展示输出的话，这里可以作为依据 取消展示
			select {
			// 检测到ctx.Done()之后停止读取
			case <-ctx.Done():
				if ctx.Err() != nil {
					if ctx.Err().Error() != "context canceled" {
						fmt.Printf("程序出现错误: %q", ctx.Err())
					}
				} else {
					fmt.Println("程序被终止")
				}
				return
			default:
				readByte, err := reader.ReadBytes('\n')
				if err != nil || err == io.EOF {
					return
				}
				fmt.Print("b => ", string(readByte))
				readString, err := reader.ReadString('\n')
				if err != nil || err == io.EOF {
					return
				}
				fmt.Print("s => ", readString)
			}
		}
	}(&wg)
	err = c.Start()
	wg.Wait()
	return err
}
