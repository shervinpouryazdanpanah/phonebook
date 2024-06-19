package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func Set() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatal("unable to decode into struct")
	}
}
