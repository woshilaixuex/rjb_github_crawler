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

func CreatModel(db *gorm.DB) {
	db.AutoMigrate(&MemberInformation{})
	db.AutoMigrate(&Information{})
}
func SaveMemberInformation(db *gorm.DB, memberInformation MemberInformation) uint {
	isSave := false
	result := db.Where("name = ?", memberInformation.Name).FirstOrCreate(&memberInformation, MemberInformation{Name: memberInformation.Name})
	if result.Error == nil && result.RowsAffected > 0 {
		return 1
	}
	for _, information := range memberInformation.Information {
		information.MemberInformationID = memberInformation.ID
		var existingInformation Information
		db.Where("data = ?", information.Data).FirstOrCreate(&existingInformation, information)
		if result.Error == nil && result.RowsAffected > 0 {
			isSave = true
		}
	}
	if isSave == true {
		return 2
	}
	return 3
}
