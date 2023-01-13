# Golang Session Manager


## 项目说明
一个Go语言Session管理器



## 初始化
```go
import "github.com/tangzhangming/session"

session.Start()
```



## 可用方法
| 方法                           |说明|
|------------------------------| ------------ |
| session.ID()                         | 返回 session_id 的值 |
| session.Name()                       | 返回 session_id 在 cookie 中的名称  |
| session.Get("user_id")               | 获得 session 某一项的值 |
| session.DefaultGet("type_id", "100") | 获得 session 某一项的值 如果不存在则返回默认值 |
| session.All()                        | 获取所有 session (map[string]string)  |
| session.Set("user_id", "10000")      | 存储一项数据到 session |
| session.Del("user_id")               | 删除 session 中一个条目  |
| session.Clear()                      | 删除 session 中所有数据  |



## 调用实例
```go
import "github.com/tangzhangming/session"

user_id := session.Get("user_id")
if user_id=="" {
	session.Get("user_id", "10001")
}

```