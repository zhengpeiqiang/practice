package customFunc_test

import (
	"fmt"
	"origin/customFunc"
	"testing"
)

func TestPrintFileStruct(t *testing.T) {
	result := ""
	//path := "C:\\Users\\cyz\\Desktop\\git-project\\practice\\go_project\\auto_test_case\\test"
	path := "E:\\1v1\\customer"
	customFunc.PrintFileStruct(path, "",1,&result)
	fmt.Println(result)
}