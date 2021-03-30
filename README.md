# Modeltools
#### GO语言连接Mysql生成对应的model，包括对应字段类型、注释等。生成基础的结构体，不局限于某一个ORM。
 
  **源码码地址---------**
  ##### github：[https://github.com/zhengpeiqiang/sqlGotomodel](https://github.com/zhengpeiqiang/sqlGotomodel)
  ##### 码云：[https://gitee.com/zhengpeiqiang/sqlGotomodel](https://gitee.com/zhengpeiqiang/sqlGotomodel)
 
 

 **生成示例---------**

```go 
  package models

  // 管理员表
  type AdminInfo struct {
  	Id int `json:"id"` 
  	Name string `json:"name"` // 姓名
  	Username string `json:"username"` // 用户名 
  	Password string `json:"password"` // 密码
  	RoleInfoId int `json:"role_info_id"` // 角色ID
  	Status int8 `json:"status"` // -1删除，0正常，1禁用
  }
```
