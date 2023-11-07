package config

import (
	"fmt"
	"log"
	"os"

	godotenv "github.com/joho/godotenv"
)

type (
	Config struct {
		App App
		Db  Db
	}

	App struct {
		Name  string
		Url   string
		Stage string
	}

	Db struct {
		Url string
	}
)

func LoadConfig(path string) Config {
	fmt.Println(path)
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		App: App{
			Name:  os.Getenv("APP_NAME"),
			Url:   os.Getenv("APP_URL"),
			Stage: os.Getenv("APP_STAGE"),
		},
		Db: Db{
			Url: os.Getenv("DB_URL"),
		},
	}

}
