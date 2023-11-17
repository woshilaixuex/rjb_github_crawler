package database

import (
	"gorm.io/gorm"
	"time"
)

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

func CreatModel(db *gorm.DB) {
	db.AutoMigrate(&MemberInformation{})
	db.AutoMigrate(&Information{})
	db.AutoMigrate(&TableUser{})
}
func SaveMemberInformation(db *gorm.DB, memberInformation MemberInformation) (uint, []Information) {
	isSave := false
	result := db.Where("name = ?", memberInformation.Name).FirstOrCreate(&memberInformation, MemberInformation{Name: memberInformation.Name})
	if result.Error == nil && result.RowsAffected > 0 {
		return 1, nil
	}
	var saveInformations []Information
	for _, information := range memberInformation.Information {
		information.MemberInformationID = memberInformation.ID
		var existingInformation Information
		result := db.Where("data = ?", information.Data).FirstOrCreate(&existingInformation, information)
		if result.Error == nil && result.RowsAffected > 0 {
			saveInformations = append(saveInformations, information)
			isSave = true
		}
	}
	if isSave == true {
		println(result.RowsAffected)
		return 2, saveInformations
	}
	return 3, nil
}
func SaveTableUser(db *gorm.DB, tbs []TableUser) {
	for _, tb := range tbs {
		db.Where("name = ?", tb.Name).FirstOrCreate(&TableUser{}, tb)
	}
}
func SaveOneTableUser(db *gorm.DB, tbs TableUser) {
	db.Where("name = ?", tbs.Name).FirstOrCreate(&TableUser{}, tbs)
}
func SelectTableUser(db *gorm.DB) []TableUser {
	var tables []TableUser
	result := db.Find(&tables)
	if result.Error != nil {

	}
	return tables
}
func SelectOneTableUser(db *gorm.DB, name string) TableUser {
	var table TableUser
	db.Where("name = ?", name).First(&table)
	return table
}
