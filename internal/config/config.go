package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"tgBotEcho/internal/model"
)

var File = &model.Config{}

func init() {
	// Загрузка файла .env
	if err := godotenv.Load("./config/echo/.env"); err != nil {
		panic(err)
	}

	err := envconfig.Process("", File)
	if err != nil {
		panic(err)
	}
	log.Println("Загруженные параметры: \n", File)

}
