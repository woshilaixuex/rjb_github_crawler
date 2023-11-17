package service

import (
	"context"
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
	"gopa/database"
	"gorm.io/gorm"
)

var (
	app_token         = "HcCtbZnBya5QOAsVHR3cIbvKnBF"
	table_id          = "tblRSjhUxroSZs7A"
	user_access_token = "u-fY2EDlClN6HbXLc59Oo214k07l.5h4jbg0w0hhi80cQX"
)

// 提交服务
func PostSever(client *lark.Client, postInformation database.MemberInformation) (string, error) {
	//resp, err := GetList(client)
	//if err != nil {
	//	return "", err
	//}
	return "", nil
}
func AddDataTable(client *lark.Client, name string) (*database.TableUser, error) {
	// 创建请求对象
	req := larkbitable.NewCreateAppTableReqBuilder().
		AppToken(app_token).
		Body(larkbitable.NewCreateAppTableReqBodyBuilder().
			Table(larkbitable.NewReqTableBuilder().
				Name(name + `github提交记录`).
				DefaultViewName(name).
				Fields([]*larkbitable.AppTableCreateHeader{
					larkbitable.NewAppTableCreateHeaderBuilder().
						FieldName(`成员`).
						Type(1).
						FieldName("信息").
						Type(1).
						FieldName("时间").
						Type(5).
						FieldName("github链接").
						Type(15).
						Build(),
				}).
				Build()).
			Build()).
		Build()
	// 发起请求
	resp, err := client.Bitable.AppTable.Create(context.Background(), req, larkcore.WithUserAccessToken(user_access_token))
	// 处理错误
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
	}
	table := database.TableUser{
		TableId: *resp.Data.TableId,
		Name:    name,
	}
	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
	return &table, nil
}

// 加上Table(用不上，你想写可以写)
//func AddTable(client *lark.Client) {
//	req := larkbitable.NewCreateAppTableReqBuilder().
//		AppToken(app_token).
//		Body(larkbitable.NewCreateAppTableReqBodyBuilder().
//			Table(larkbitable.NewReqTableBuilder().
//				Name(`数据表名称2`).
//				DefaultViewName(`默认的表格视图`).
//				Fields([]*larkbitable.AppTableCreateHeader{
//					larkbitable.NewAppTableCreateHeaderBuilder().
//						FieldName(`多行文本`).
//						FieldName(``).
//						Type(1).
//						Build(),
//				}).
//				Build()).
//			Build()).
//		Build()
//	// 发起请求
//	resp, err := client.Bitable.AppTable.Create(context.Background(), req, larkcore.WithUserAccessToken("u-cIwzjSDW9bxWmVNK1S4.3.ggjRf514V3rww0glG02FWu"))
//	// 处理错误
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	// 服务端错误处理
//	if !resp.Success() {
//		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
//		return
//	}
//	// 业务处理
//	fmt.Println(larkcore.Prettify(resp))
//}

// 获取数据表信息
func GetDataTables(client *lark.Client) ([]database.TableUser, error) {
	// 创建请求对象
	req := larkbitable.NewListAppTableReqBuilder().
		AppToken(app_token).
		Build()

	// 发起请求
	resp, err := client.Bitable.AppTable.List(context.Background(), req, larkcore.WithUserAccessToken(user_access_token))
	// 处理错误
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return nil, err
	}
	var tables []database.TableUser
	for _, item := range resp.Data.Items {
		tableUser := database.TableUser{
			TableId: *item.TableId,
			Name:    *item.Name,
		}
		tables = append(tables, tableUser)
	}
	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
	return tables, nil
}

// 添加视图
//func AddView(client *lark.Client) {
//	// 创建请求对象
//	req := larkbitable.NewCreateAppTableViewReqBuilder().
//		AppToken(app_token).
//		TableId(table_id).
//		ReqView(larkbitable.NewReqViewBuilder().
//			ViewName(`表格视图`).
//			ViewType(`grid`).
//			Build()).
//		Build()
//	// 发起请求
//	resp, err := client.Bitable.AppTableView.Create(context.Background(), req, larkcore.WithUserAccessToken(user_access_token))
//	// 处理错误
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	// 服务端错误处理
//	if !resp.Success() {
//		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
//		return
//	}
//	// 业务处理
//	fmt.Println(larkcore.Prettify(resp))
//}

// 添加列
func AddList(db *gorm.DB, client *lark.Client, posts database.MemberInformation) {
	var records []*larkbitable.AppTableRecord
	for _, info := range posts.Information {
		record := larkbitable.NewAppTableRecordBuilder().
			Fields(map[string]interface{}{
				"成员":       posts.Name,
				"信息":       info.Message,
				"日期":       info.Data,
				"github链接": info.Url,
			}).
			Build()
		records = append(records, record)
	}
	req := larkbitable.NewBatchCreateAppTableRecordReqBuilder().
		AppToken(app_token).
		TableId(database.SelectOneTableUser(db, posts.Name).TableId).
		Body(larkbitable.NewBatchCreateAppTableRecordReqBodyBuilder().
			Records(records).
			Build()).
		Build()
	// 发起请求
	resp, err := client.Bitable.AppTableRecord.BatchCreate(context.Background(), req, larkcore.WithUserAccessToken(user_access_token))
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

// 获取列
func GetList(client *lark.Client) (*larkbitable.ListAppTableRecordResp, error) {
	req := larkbitable.NewListAppTableRecordReqBuilder().
		AppToken(app_token).
		TableId(table_id).
		Build()
	// 发起请求
	resp, err := client.Bitable.AppTableRecord.List(context.Background(), req, larkcore.WithUserAccessToken(user_access_token))
	// 处理错误
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return nil, err
	}
	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
	return resp, err
}
