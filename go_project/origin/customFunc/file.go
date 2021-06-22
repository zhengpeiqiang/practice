package customFunc

import (
	"errors"
	"io/ioutil"
	"os"
)

/*
* action 打印文件夹结构层次
* param path 文件夹地址
* param indent 每行前缀；默认不需要，主要用于递归使用
* param dirFirst 是否打印文件夹在前；1 是，0 否
* return string
 */
func PrintFileStruct(path string, indent string, dirFirst int, result *string) {
	if path == "" {
		exPath, err := os.Getwd() // 获取程序执行的当前路径
		if err != nil {
			WriteLog("", "", nil, errors.New("路径错误"))
		}
		path = exPath
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		WriteLog("", "", err, errors.New("read file path error"))
	}
	// 忽略以 . 开头的文件
	for i := 0; i < len(files); i++ {
		if files[i].Name()[0] == '.' {
			files = append(files[:i], files[i+1:]...)
		}
	}
	dirs := make([]string, 0)

	// 先打印文件
	for _, fi := range files {
		if !fi.IsDir() {
			dirs = append(dirs, fi.Name())
		}
	}
	lenFile := len(dirs)

	// 再打印文件夹
	for _, fi := range files {
		if fi.IsDir() {
			dirs = append(dirs, fi.Name())
		}
	}

	if dirFirst == 1 {
		for i := len(dirs) - 1; i >= 0; i-- {
			if i == 0 {
				*result = *result + indent + "└── " + dirs[i] + "\n"
				if i >= lenFile {
					PrintFileStruct(path+"\\"+dirs[i], indent+"   ", dirFirst, result)
				}
			} else {
				*result = *result + indent + "├── " + dirs[i] + "\n"
				if i >= lenFile {
					PrintFileStruct(path+"\\"+dirs[i], indent+"│  ", dirFirst, result)
				}
			}
		}
	} else {
		for i := 0; i < len(dirs); i++ {
			if i == len(dirs)-1 {
				*result = *result + indent + "└── " + dirs[i] + "\n"
				if i >= lenFile {
					PrintFileStruct(path+"\\"+dirs[i], indent+"   ", dirFirst, result)
				}
			} else {
				*result = *result + indent + "├── " + dirs[i] + "\n"
				if i >= lenFile {
					PrintFileStruct(path+"\\"+dirs[i], indent+"│  ", dirFirst, result)
				}
			}
		}
	}
}
