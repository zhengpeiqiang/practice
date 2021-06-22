# Test-路西法

### 文件分层
```
├── build           生成执行文件目录
├── cache           缓存
│   └── cacheList   缓存链表
│   └── cacheMap    缓存字典
├── conf            配置文件夹
├── customFunc      自定义公共方法
├── initFunc        初始化
├── logs            日志
├── middleware      中间件
├── modules         公共的方法模块
├── public          公共文件
├── router          路由层
│   └── controller  控制器
│   └── models      数据库模型
│   └── service     数据库模型方法
├── timeTask        定时任务
├── go.mod          包
└── main.go         入口文件
└── README.md
```

访问url：https://127.0.0.1:9999/test/welcome