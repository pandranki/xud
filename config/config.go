package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Viper *viper.Viper

//InitConfig ...
func InitConfig() *viper.Viper {
	viper.SetConfigName("xud") //name of config file (without extension)
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("~/.xud/")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	
	if err != nil {
		e, ok := err.(viper.ConfigParseError)
		if ok {
			log.Fatalf("Error parsing config file: %v", e)
		}
	}
	
	Viper = viper.GetViper()
	return Viper
}
