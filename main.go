package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"go_training/config"
	"go_training/initializer"
	"go_training/web/handler"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//start()
	conf := config.NewConfig()
	infras := initializer.InitInfras(&conf)
	infras.EmailSender.SendEmail("hoge@example.com")
}

func start() {
	conf := config.NewConfig()
	db, err := gorm.Open("mysql", conf.DB.GetSettingStr())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	handler.InitServer(conf, db)
}
