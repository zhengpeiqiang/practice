package task

import (
	"fmt"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/util/gconv"
	"time"
	"use_granphviz/pressure"
)

const tt = 1 * time.Second

var TaskList *gmap.Map

type TaskListContent struct {
	RunNum        int
	ConcurrentNum int
	Type          string
	Maps          map[string]interface{}
}

var numPressure = 0
var numChannel = 0

func InitTask() {
	TaskList = gmap.New()
	gtimer.AddSingleton(tt, func() {
		types := ""
		key := ""
		runNum := 0
		concurrentNum := 0
		maps := make(map[string]interface{}, 0)
		TaskList.Iterator(func(k interface{}, v interface{}) bool {
			taskListContent := v.(TaskListContent)
			types = taskListContent.Type
			key = k.(string)
			runNum = taskListContent.RunNum
			switch taskListContent.Type {
			case "pressure":
				if taskListContent.RunNum > 0 {
					pressure.Pressure(
						taskListContent.ConcurrentNum,
						0, "oR52qv4HG8K1MyFdGhXaH6Qjqz7A", 2,
						"PQ", "https://thirdwx.qlogo.cn/mmopen/vi_32/j7SnZnaxybGgsYNl4YNIpJibAH6JZdxImoUiaAWz2paK33BGYu9EEMsBvBxuZ0cvZop7YCZfhpamzN6ZCURfuwow/132", "127317",
						"119217")
					runNum--
					numPressure = numPressure + taskListContent.ConcurrentNum
				}
				concurrentNum = taskListContent.ConcurrentNum
				break
			case "channel":
				if taskListContent.RunNum > 0 {
					pressure.Channel(taskListContent.ConcurrentNum, gconv.Int(taskListContent.Maps["channel_type"]))
					runNum--
					numChannel = numChannel + taskListContent.ConcurrentNum
				}
				concurrentNum = taskListContent.ConcurrentNum
				maps = taskListContent.Maps
				break
			}
			return true
		})
		if key != "" && types == "pressure" {
			fmt.Println("压测结束 错误数 => ", gconv.String(pressure.ErrorNum), " num => ", numPressure, " time => ", time.Now().Format("2006-01-02 15:04:05"))
			if runNum <= 0 {
				TaskList.Remove(key)
			} else {
				TaskList.Set(key, TaskListContent{
					RunNum:        runNum,
					ConcurrentNum: concurrentNum,
					Type:          "pressure",
				})
			}
		}
		if key != "" && types == "channel" {
			fmt.Println("压测结束 错误数 => ", gconv.String(pressure.ErrorNum), " num => ", numChannel, " time => ", time.Now().Format("2006-01-02 15:04:05"))
			if runNum <= 0 {
				TaskList.Remove(key)
			} else {
				TaskList.Set(key, TaskListContent{
					RunNum:        runNum,
					ConcurrentNum: concurrentNum,
					Type:          "channel",
					Maps:          maps,
				})
			}
		}
	})
}
