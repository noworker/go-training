package config

import (
	"fmt"
	"os"
)

const DbArgs string = "%s:%s@tcp(%s)/%s?parseTime=true&loc=%s"

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Host     string
	UserName string
	Password string
	Name     string
	Location string
}

func NewConfig() Config {
	conf := Config{}
	conf.DB.Host = os.Getenv("SYSTEMS_DB_HOST")
	conf.DB.Name = os.Getenv("SYSTEMS_DB_NAME")
	conf.DB.UserName = os.Getenv("SYSTEMS_DB_USER")
	conf.DB.Password = os.Getenv("SYSTEMS_DB_PASSWORD")
	conf.DB.Location = "Asia%2FTokyo"
	return conf
}

func (dbConfig *DBConfig) GetSettingStr() string {
	return fmt.Sprintf(DbArgs,
		dbConfig.UserName,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Name,
		dbConfig.Location,
	)
}
