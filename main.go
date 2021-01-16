package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"go_training/config"
	"go_training/infrastructure/jw_token"
	"go_training/web/handler"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print(err.Error())
		log.Fatal("Error loading .env file")
	}
	start()
}

func start() {
	conf := config.NewConfig()
	jw_token.NewKeyGenerator("./").Generate(conf)
	db, err := gorm.Open("mysql", conf.DB.GetSettingStr())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	handler.InitServer(conf, db)

}
