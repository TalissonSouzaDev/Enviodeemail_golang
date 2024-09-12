package config

import "github.com/spf13/viper"

type config struct {
	EmailHost     string `mapstructure:"EMAIL_HOST"`
	EmailPort     string `mapstructure:"EMAIL_PORT"`
	UserEmail     string `mapstructure:"EMAIL_USER"`
	PasswordEmail string `mapstructure:"EMAIL_PASSWORD"`
}

func LoadConfig() (*config, error) {
	viper.SetConfigFile("cmd/.env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var cfg config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg, nil
}
