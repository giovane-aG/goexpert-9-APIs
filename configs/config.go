package configs

import (
	"gihub.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPass        string
	DBName        string
	WebServerPort string
	JWTSecret     string
	JWTExpiresIn  int
	TokenAuth     *jwtauth.JWTAuth
}

var cfg *conf

func LoadConfig(path string) *conf {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
