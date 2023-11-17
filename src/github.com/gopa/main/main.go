package main

import "gopa/service"

// SDK 使用文档：https://github.com/larksuite/oapi-sdk-go/tree/v3_main
func main() {
	service.DBServer()
	//db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gopa_db?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	//database.CreatModel(db)
}
