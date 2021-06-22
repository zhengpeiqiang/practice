package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

func main()  {
	srcByte, err := ioutil.ReadFile(`C:\Users\cyz\Desktop\git-project\fyne-use\src\picture\fyne.png`)
	if err != nil {
		log.Fatal(err)
	}

	res := base64.StdEncoding.EncodeToString(srcByte)

	fmt.Println(res)

}
