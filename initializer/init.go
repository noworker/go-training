package initializer

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go_training/config"
)

func Init() {
	conf := config.NewConfig()
	db, err := gorm.Open("mysql", conf.DB.GetSettingStr())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	InitRepositories(db)
}
