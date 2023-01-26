package test

import (
	"app/models"
	"app/utils"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestGorm() {

	// 迁移 schema
	utils.DB.AutoMigrate(&models.AppUser{})

}
