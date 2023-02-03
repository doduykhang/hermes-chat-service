package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Server struct {
	Port string `mapstructure:"PORT"`
}

type Redis struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
	Password string `mapstructure:"PASSWORD"`
}

type DB struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
	User string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"` 
	Name string `mapstructure:"NAME"`
}

type RabbitMQ struct {
	Protocol string `mapstructure:"PROTOCOL"` 
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
	User string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"` 
	VHost string `mapstructure:"VHOST"` 
}

type Config struct {
	Server Server `mapstructue:"SERVER"`
	Redis Redis `mapstructure:"REDIS"`
	DB DB `mapstructure:"DB"`
	RabbitMQ RabbitMQ `mapstructure:"RABBITMQ"`
}

func LoadEnv(path string) (*Config) {
	replacer := strings.NewReplacer(".", "_")
    	viper.SetEnvKeyReplacer(replacer)
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading env, %s\n", err)
	}

	var env Config 
	err = viper.Unmarshal(&env)
	return &env
}
