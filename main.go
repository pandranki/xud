package main

import (
	"runtime"
	"github.com/indxcrypto/xud/config"
	database "github.com/indxcrypto/xud/db"
	"github.com/indxcrypto/xud/lib/core"
)

func main() {
	println("Exchange Union Daemon")
	runtime.GOMAXPROCS(runtime.NumCPU())
	config.Viper = config.InitConfig() //Read Config
	viper := config.Viper                      //expose viper
	database.DBCon() //Connect to database
	db := database.Db
	database.SchemaAutoMigrate(db)
	xud := core.NewXUD(db,viper)
	xud.Start() //Starts XUD
	xud.Shutdown() //Stops XUD
}
