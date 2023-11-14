package main

import (
	"context"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

func AddList(client *lark.Client) {
	req := larkbitable.NewCreateAppTableRecordReqBuilder().
		AppToken(`HcCtbZnBya5QOAsVHR3cIbvKnBF`).
		TableId(`tblRSjhUxroSZs7A`).
		AppTableRecord(larkbitable.NewAppTableRecordBuilder().
			Fields(map[string]interface{}{`commit`: `1`, `成员`: `小莫`, `日期`: `2022-1-2`, `github链接`: `https://open.feishu.cn/document/server-docs/docs/bitable-v1/app-table-record/create`}).
			Build()).
		Build()

	// 发起请求
	resp, err := client.Bitable.AppTableRecord.Create(context.Background(), req, larkcore.WithUserAccessToken("u-cIwzjSDW9bxWmVNK1S4.3.ggjRf514V3rww0glG02FWu"))

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return
	}

	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
}
func AddTable(client *lark.Client) {
	req := larkbitable.NewCreateAppTableReqBuilder().
		AppToken(`HcCtbZnBya5QOAsVHR3cIbvKnBF`).
		Body(larkbitable.NewCreateAppTableReqBodyBuilder().
			Table(larkbitable.NewReqTableBuilder().
				Name(`数据表名称`).
				DefaultViewName(`默认的表格视图`).
				Fields([]*larkbitable.AppTableCreateHeader{
					larkbitable.NewAppTableCreateHeaderBuilder().
						FieldName(`多行文本`).
						Type(1).
						Build(),
				}).
				Build()).
			Build()).
		Build()

	// 发起请求
	resp, err := client.Bitable.AppTable.Create(context.Background(), req, larkcore.WithUserAccessToken("u-cIwzjSDW9bxWmVNK1S4.3.ggjRf514V3rww0glG02FWu"))

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return
	}

	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
}
