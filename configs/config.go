package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var Config *config

type (
	config struct {
		App app `yaml:"app"`
		DB  db  `yaml:"db"`
	}

	app struct {
		Name        string `yaml:"name"`
		Environment string `yaml:"environment"`
		Url         string `yaml:"url"`
		Key         string `yaml:"key"`
	}

	db struct {
		Postgres postgres `yaml:"postgres"`
	}

	postgres struct {
		Connection string `yaml:"connection"`
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		Database   string `yaml:"database"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		SSLMode    string `yaml:"sslmode"`
	}
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("не удалось спарсить конфиг файл. Error: %s", err)

		log.Fatal(err)
	}

	Config = &config{
		App: app{
			Name:        viper.GetString("app.name"),
			Environment: viper.GetString("app.environment"),
			Url:         viper.GetString("app.url"),
			Key:         viper.GetString("app.key"),
		},
		DB: db{
			Postgres: postgres{
				Connection: viper.GetString("db.postgres.connection"),
				Host:       viper.GetString("db.postgres.host"),
				Port:       viper.GetString("db.postgres.port"),
				Database:   viper.GetString("db.postgres.database"),
				Username:   viper.GetString("db.postgres.username"),
				Password:   viper.GetString("db.postgres.password"),
				SSLMode:    viper.GetString("db.postgres.sslmode"),
			},
		},
	}
}
