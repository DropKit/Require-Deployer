package main 

import (
	"log"

	"github.com/DropKit/Require-Deployer/contracts/authority"
	"github.com/DropKit/Require-Deployer/contracts/metaTable"
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
	metaTableAddr, _ := metaTable.Deploy()
	log.Print("MetaTable Address: " + metaTableAddr)

	authorityAddr, _ := authority.Deploy()
	log.Print("Auhority Address: " + authorityAddr)
}