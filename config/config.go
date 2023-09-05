package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DB DataBase
}

type DataBase struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"passwd"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Type     string `mapstructure:"type"`
}

func LoadConfig(path string) Config {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	var config Config
	godotenv.Load(".env")

	if str, ok := os.LookupEnv("DB_HOST"); ok {
		config.DB.Host = str
	} else {
		config.DB.Host = viper.GetString("DB.HOST")
	}
	if str, ok := os.LookupEnv("DB_NAME"); ok {
		config.DB.Name = str
	} else {
		config.DB.Name = viper.GetString("DB.NAME")
	}
	if str, ok := os.LookupEnv("DB_PASSWORD"); ok {
		config.DB.Password = str
	} else {
		config.DB.Password = viper.GetString("DB.PASSWORD")
	}
	if str, ok := os.LookupEnv("DB_PORT"); ok {
		config.DB.Port, _ = strconv.Atoi(str)
	} else {
		config.DB.Port = viper.GetInt("DB.PORT")
	}
	if str, ok := os.LookupEnv("DB_TYPE"); ok {
		config.DB.Type = str
	} else {
		config.DB.Type = viper.GetString("DB.type")
	}
	if str, ok := os.LookupEnv("DB_USER"); ok {
		config.DB.User = str
	} else {
		config.DB.User = viper.GetString("DB.USER")
	}
	return config
}
