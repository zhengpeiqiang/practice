package customFunc

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"io/ioutil"
)

func WriteLog(path string, filename string, data interface{}, errMsg error) {
	err := ioutil.WriteFile(path+filename, gconv.Bytes(data), 0777)
	if err != nil {
		fmt.Println("File reading error", err)
	}
}

func ReadLog(path string, filename string) {
	data, err := ioutil.ReadFile(path + filename)
	if err != nil {
		fmt.Println("File reading error", err)
	}
	fmt.Println(string(data))
}
