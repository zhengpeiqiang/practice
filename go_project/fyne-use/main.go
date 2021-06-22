package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"path/filepath"
)

func main() {
	//`C:\Users\cyz\Desktop\git-project\fyne-use\src\picture\lclcf001.jpg`
	a := app.NewWithID("io.fyne.zpq")
	a.SetIcon(theme.FyneLogo())
	w := a.NewWindow("智能小工具")
	//topWindow = w

	//widget.NewList()

	var chosen fyne.URIReadCloser
	var openErr error
	d := dialog.NewFileOpen(func(file fyne.URIReadCloser, err error) {
		chosen = file
		openErr = err
	}, w)
	testDataPath, _ := filepath.Abs("src")
	testData := storage.NewFileURI(testDataPath)
	dir, err := storage.ListerForURI(testData)
	if err != nil {
		fmt.Println(err)
	}
	d.SetLocation(dir)
	d.Show()

	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}
