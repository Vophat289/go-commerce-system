package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/") //path to config

	viper.SetConfigName("production") //ten file
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read config %w \n", err))
	}
	//write server config
	fmt.Println("Server Port::", viper.GetInt("server.port"))
	fmt.Println("Security Key Port::", viper.GetString("security.jwt.key"))

}
