package config

import( "github.com/spf13/viper"
log "github.com/sirupsen/logrus")

var Viper *viper.Viper

//InitConfig ...
func InitConfig() *viper.Viper {
	viper.SetConfigName("xud") // name of config file (without extension)
	viper.AddConfigPath(".")
	viper.AddConfigPath("~/.xud/")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		e, ok := err.(viper.ConfigParseError)
		if ok {
			log.Fatalf("Error parsing config file: %v", e)
		}
		log.Debugf("No config file used")
	} else {
		log.Debugf("Using config file: %v", viper.ConfigFileUsed())
	}
	Viper = viper.GetViper()
	return Viper
}
