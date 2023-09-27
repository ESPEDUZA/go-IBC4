package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
}

func SetupModels(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return
	}
}
