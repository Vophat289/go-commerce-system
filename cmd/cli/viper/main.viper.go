package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct{
	Server struct {
		Port int `mapstructure:"port"`
	}`mapstructure:"server"`	
	Database []struct{
		User string `mapstruct:"user"`
		Password string `mapstruct:"password"`
		Host string `mapstruct:"host"`
	}`mapstructure:"databases"`
}

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

	//config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode config %v", err)
	}

	fmt.Println("Config Port::", config.Server.port)
	

}
