package main

import (
	"log"

	dropkit "github.com/DropKit/Require-Deployer/contracts/dropkit"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		print(err)
	}
}

func main() {
	dropkitAddress, _ := dropkit.Deploy()
	log.Print("DropKit contract address: " + dropkitAddress)
}
