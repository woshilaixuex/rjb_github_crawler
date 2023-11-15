package main

import "fmt"

// SDK 使用文档：https://github.com/larksuite/oapi-sdk-go/tree/v3_main
func main() {
	// 创建 Client
	//client := lark.NewClient("cli_a5d9d68ff4f8500c", "9TeR7xstUEuiuMvHMbFQffqg8AtXlXsC",
	//	lark.WithLogLevel(larkcore.LogLevelDebug),
	//	lark.WithReqTimeout(3*time.Second),
	//	lark.WithHttpClient(http.DefaultClient),
	//)
	postInformation, err := GetCommit()
	if err != nil {
		panic(err)
	}
	fmt.Println(postInformation)
}
