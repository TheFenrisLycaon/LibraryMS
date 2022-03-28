package models

import (
	"github.com/TheFenrisLycaon/LibraryMS/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name      string `gorm:"" json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Price     int    `json:"price"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
