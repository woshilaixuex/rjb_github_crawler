package service

import (
	"gopa/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:123456@tcp(127.0.0.1:3306)/gopa_db?charset=utf8mb4&parseTime=True&loc=Local"

func DBServer(memberInformations []database.MemberInformation) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//database.CreatModel(db)
	for _, memberInformation := range memberInformations {
		database.SaveMemberInformation(db, memberInformation)
	}
}
