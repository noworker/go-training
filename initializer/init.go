package initializer

import "github.com/jinzhu/gorm"

func Init(db *gorm.DB) {
	InitRepositories(db)
}