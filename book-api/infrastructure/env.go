package infrastructure

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`
	DBUsername  string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBName      string `mapstructure:"DB_NAME"`
	LogOutput   string `mapstructure:"LOG_OUTPUT"`
}

var globalEnv = Env{}

func GetEnv() Env {
	return globalEnv
}

func NewEnv() Env {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot read cofiguration")
	}

	err = viper.Unmarshal(&globalEnv)
	if err != nil {
		log.Fatal("environment cant be loaded: ", err)
	}

	log.Printf("%#v \n", &globalEnv)
	return globalEnv
}
