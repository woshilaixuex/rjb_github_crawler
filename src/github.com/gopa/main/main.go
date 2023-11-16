package main

import (
	"gopa/service"
)

var (
	appId     = "cli_a5d9d68ff4f8500c"
	appSecret = "9TeR7xstUEuiuMvHMbFQffqg8AtXlXsC"
)

// SDK 使用文档：https://github.com/larksuite/oapi-sdk-go/tree/v3_main
func main() {
	//创建 Client
	//client := lark.NewClient(appId, appSecret,
	//	lark.WithLogLevel(larkcore.LogLevelDebug),
	//	lark.WithReqTimeout(5*time.Second),
	//	lark.WithHttpClient(http.DefaultClient),
	//)
	postInformation, err := GetCommit()
	if err != nil {
		panic(err)
	}
	//AddList(client, postInformation[0])
	//AddView(client)
	//GetList(client)
	service.DBServer(postInformation)
}
