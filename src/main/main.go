package main

import (
	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"net/http"
	"time"
)

// SDK 使用文档：https://github.com/larksuite/oapi-sdk-go/tree/v3_main
func main() {
	// 创建 Client
	client := lark.NewClient("cli_a5d9d68ff4f8500c", "9TeR7xstUEuiuMvHMbFQffqg8AtXlXsC",
		lark.WithLogLevel(larkcore.LogLevelDebug),
		lark.WithReqTimeout(3*time.Second),
		lark.WithHttpClient(http.DefaultClient),
	)
	AddList(client)
}
