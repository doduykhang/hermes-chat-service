package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	RedisHost string `mapstructure:"REDIS_HOST"`
	RedisPort string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"` 
	DBName string `mapstructure:"DB_NAME"`
}

func LoadEnv(path string) (*Env) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading env, %s\n", err)
	}

	var env Env
	err = viper.Unmarshal(&env)
	return &env
}
