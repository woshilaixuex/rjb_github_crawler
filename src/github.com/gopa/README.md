# 软件部Github提交记录爬取

##  核心

###  业务逻辑

server里是对这个项目的整合步骤分为

1. 注册组件
2. 拉取多维表格中的的数据与数据库比对实现数据同步
3. 爬取用户的Commit信息
4. 与数据库中数据对比判断是否有未记录数据（有就想飞书发送添加数据请求）

crawler_server(爬取)和post_server(通过飞书api操作多维表格)，

##  模型

###  爬取并转换json(Commit数据)

```go
type Commit struct {
	SHA    string `json:"sha"`
	Commit struct {
        //这里提交人的用户信息
		Author struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Date  string `json:"date"`
		} `json:"author"`
		Message string `json:"message"`
	} `json:"commit"`
    //仓库所属人的信息
	Author struct {
		Login string `json:"login"`
	} `json:"author"`
}
```

###  用户

```go
type MemberInformation struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Information []Information `gorm:"foreignKey:MemberInformationID"`
}

type Information struct {
	ID                  uint `gorm:"primaryKey"`
	MemberInformationID uint
	Message             string
	Data                time.Time
	Url                 string
}
type TableUser struct {
	TableId string `json:"TableId"`
	Name    string `json:"Name"`
}
```

##  服务端写完了来个客户端!!!

