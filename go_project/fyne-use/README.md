<p align="center">
<img alt="" src="https://raw.githubusercontent.com/zhengpeiqiang/LCLCF_origin/master/zhengpeiqiang.png" style="border-radius:50%;margin: 0 auto;" width="20%" />
</p>

<h1 align="center">
fyne use
</h1>

### organization

<a href="https://github.com/zhengpeiqiang"><img alt="" src="https://raw.githubusercontent.com/zhengpeiqiang/LCLCF_origin/master/LCLCF_circle.png" style="width:60px;height:60px;margin: 0 auto;" width="8%" /></a>
**LCLCF**

### use fyne

```
fyne 实现

0.如果是 window 的话，安装 gcc 
    可以参考 https://www.shangmayuan.com/a/20a48fa4e1c34b63a98ce8f9.html 对照操作即可

1.引入 fyne，实际作用不明
    $ go get fyne.io/fyne/v2

2.尝试 fyne 的官方 demo
    $ go get fyne.io/fyne/v2/cmd/fyne_demo
    $ fyne_demo

3.写一个自己的 app ，代码如下
    package main
    
    import (
    	"fyne.io/fyne/v2"
    	"fyne.io/fyne/v2/app"
    	"fyne.io/fyne/v2/theme"
    )
    
    func main() {
    	//`C:\Users\cyz\Desktop\git-project\fyne-use\src\picture\lclcf001.jpg`
    	a := app.NewWithID("io.fyne.zpq")
    	a.SetIcon(theme.FyneLogo())
    	w := a.NewWindow("智能小工具")
    	//topWindow = w
    	w.Resize(fyne.NewSize(640, 460))
    	w.ShowAndRun()
    }

4.引入 fune cmd ，用于打包
    $ go get fyne.io/fyne/cmd/fyne
    可参考文档 https://studygolang.com/articles/33444?fr=sidebar
    
```

