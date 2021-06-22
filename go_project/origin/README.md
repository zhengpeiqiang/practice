## ![avatar](https://portrait.gitee.com/uploads/avatars/namespace/2645/7936029_no_anyone_care_1623389921.png!avatar60) L C L C F

``` 
LCLCF https://gitee.com/no_anyone_care
        _______  _______  ___  _______  ___  _______
       / ___  / / ___  / /__/ /  __  / /__/ / ___  /
      / /  / / / /__/_/ ___  / / ___  ___  / /  / /
     / /  / / / ___ \  /  / / /  / / /  / / /  / /
    / /__/ / / /  / / /  / / /__/ / /  / / /  / /
   /______/ /_/  /_/ /__/ /______/ /__/ /_/  /_/

```

```
1  =>  2  =>  3  =>  4
双缝干涉
    意识 与 感知
```

```
乐高 Plugin

从 主键模块 重新理解系统本身
    使用到的主键
        C、C++、C#、PHP、JAVA、Go、Javascript
        Nginx、Apache、Tomcat
        Oracle、MySQL、SqlServer、PostgreSQL

从 功能模块 重新理解系统本身

从 操作流程 重新理解系统本身
    【用户，系统，管理，服务人员。。。】登录，有则登录并校验，没有则注册并登录
    发送消息
        例如，商城内，【A人类·用户】上架了一款产品，接收方是【上架·系统】，
        然后它（他/她/祂）的权限是把商品放到 “架子上”--它（他/她/祂）的接收框
    接收消息
        
    
```

# origin 框架

- 框架思想

每个模块及主键专注自己的业务不混淆

- 设置接口文档 controller 层的 Bind 方法的注释设置

```
// @Summary 登录
// @Description
// @Tags 用户中心
// @ID   /login
// @Accept  json
// @Produce  json
// @Param polygon body interfaceStruct.LoginInput true "body"
// @Success 200 {object} middleware.Response{data=interfaceStruct.LoginOutput} "success"
// @Router /login [post]

生成接口文档：swag init
浏览地址：http://127.0.0.1:9090/customer/swagger/index.html
```

- 代码文件结构

```
├── build           生成执行文件目录
├── cache           缓存
│   └── cacheList   缓存链表
│   └── cacheMap    缓存字典
├── conf            配置文件夹
├── customFunc      自定义公共方法
├── file            文件
│   └── music       音乐文件
├── initFunc        初始化
├── logs            日志
├── middleware      中间件
├── modules         公共的方法模块
├── public          公共文件
├── router          路由层
│   └── controller  控制器
│   └── models      数据库模型
│   └── service     数据库模型方法
│   └── view        模板
├── test            测试
├── timeTask        定时任务
├── go.mod          包
└── main.go         入口文件
└── README.md
```

- 程序启动执行流程 

```
程序在重启时，可以完全接续之前的服务，即使是在并发过程中
1.步骤是 先 读取配置信息，然后 读取上一份缓存备份，再然后 是处理request请求
```

- 程序request-response模型

```
如果request请求超时，可以约定一个超时返回，让前端在大概多久之后重新请求，且一个"原请求"累计重新请求次数不得超过某个值,超过即放弃，这些都是后端进行
请求msg ok【成功】，fail【失败】，timeout【超时】，timeout-close【超时关闭】

成功
|client             (request)->                 server|
|client         <-(response,msg=ok)             server|

失败
|client             (request)->                 server|
|client         <-(response,msg=fail)           server|

重新请求，re-request（再次请求时间）
|client             (request)->                 server|
|client <-(response,msg=timeout,re-request=1s)  server|
|client          (request,re-first)->           server|
|client <-(response,msg=timeout,re-request=1s)  server|
|client          (request,re-second)->          server|
|                       ...                           |
|                       ...                           |
|client     <-(response,msg=timeout-close)      server|
```

- 缓存策略

```
1.不允许默认数据的缓存策略，且数据不可开发人员创建
例子：查询用户数据
① 查缓存
    ② 存在，返回 => 结束
    ② 不存在
        ③ 查询DB
            ④ 存在，存缓存，返回 => 结束

2.允许默认数据的缓存策略
例子：查询用户数据
① 查缓存
    ② 存在，返回 => 结束
    ② 不存在
        ③ 查询DB
            ④ 存在，存缓存，返回 => 结束
            ④ 不存在，添加默认数据，存缓存，返回 => 结束

3.不允许默认数据的缓存策略，数据可以开发人员创建
例子：查询用户数据
初次查询：
① 查缓存
    ② 存在，返回 => 结束
    ② 不存在，返回 => 结束
创建：
① 查缓存
    ② 存在，返回 => 结束
    ② 不存在，添加数据，存缓存 => 结束
再次查询：
① 查缓存
    ② 存在，返回 => 结束
    ② 不存在，返回 => 结束

两种策略，根据情况选择使用，如果是 服务间能够通信的 建议采用第三种 
可以减少缓存被击穿的情况 
选择第二种策略也行 
第一种策略 被攻击时 可能会直接把数据库 负载拉的非常高
```

#### 打包编译

```
run.bat

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o build/customer_web main.go
```

#### 压测方式

```
安装【安装不了的自行下载 zip 源码安装】
go get -u github.com/rakyll/hey

详细指令说明看 https://github.com/rakyll/hey

执行
hey -n 300 ^
    -c 20 ^
    -m POST ^
    -T "application/x-www-form-urlencoded" ^
    -d "unionId=oR52qv5CbbgLKqc_mm5rTTofrer8&isCustomerClient=1&channelType=2" ^
    http://127.0.0.1:8888/customer/checkuser
```

#### git 指令说明

```
拉取【origin master，远程 master 分支】
git pull origin master
添加
git add .
提交
git commit -m"xxxxxxxx"
推送【origin master，远程 master 分支】
git push origin master

查看远程分支
git branch -r
查看本地分支
git branch -a

更新分支信息
git remote update origin --prune

切换分支
git checkout origin/v1.0
切换分支，且命名分支
git checkout -b v1.0 origin/v1.0
    后续拉取，推送
    git pull origin v1.0
    git push origin v1.0
```
