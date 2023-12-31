package service

import (
	"fmt"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"gopa/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var (
	appId     = "cli_a5d9d68ff4f8500c"
	appSecret = "9TeR7xstUEuiuMvHMbFQffqg8AtXlXsC"
)

const dsn = "root:123456@tcp(127.0.0.1:3306)/gopa_db?charset=utf8mb4&parseTime=True&loc=Local"

func DBServer() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database.CreatModel(db)
	client := lark.NewClient(appId, appSecret,
		lark.WithLogLevel(larkcore.LogLevelDebug),
		lark.WithReqTimeout(5*time.Second),
		lark.WithHttpClient(http.DefaultClient),
	)
	//数据同步
	cltables, err := GetDataTables(client)
	cltables = cltables[1:]
	if err != nil {
		panic(err)
	}
	database.SaveTableUser(db, cltables)

	postInformation, ero := GetCommit()
	if ero != nil {
		panic(ero)
	}
	fmt.Println(postInformation)
	for _, postinfor := range postInformation {
		action, saveInformations := database.SaveMemberInformation(db, postinfor)
		if action == 1 { //没有存过用户
			newtabls, err := AddDataTable(client, postinfor.Name)
			if err != nil {
				panic(err)
			}
			database.SaveOneTableUser(db, *newtabls)
			database.SaveMemberInformation(db, postinfor)
			AddList(db, client, postinfor)
		}
		if action == 2 {
			//这边有问题
			newpost := database.MemberInformation{
				Name:        postinfor.Name,
				Information: saveInformations,
			}
			AddList(db, client, newpost)
		}
	}
}
