package customFunc

import (
	"fmt"
	"math/rand"
)

var defaultLetters = []rune(fmt.Sprintf("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"))

/*
* action 生成拟定长度随机字符串
* param randomLength 随机值长度
* param allowedChars 重定义字符串基础字符
* return string
 */
func RandomString(randomLength int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, randomLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}