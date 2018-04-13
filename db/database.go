package db

import (
	"github.com/indxcrypto/xud/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	//Db Exported DB Connection
	Db *gorm.DB
)

// DBCon creates new database connection.
func DBCon() *gorm.DB {
	db, err := gorm.Open("sqlite3", "xud.db")
	if err != nil {
		log.Errorln(err)
		log.Fatalln("Failed to connect database.")
	}
	log.Infoln("Connected to Database Server")
	//defer db.Close()
	if viper.GetBool("db.log") {
		db.LogMode(true)
	}
	Db = db
	return Db
}

//SchemaAutoMigrate automatically migrates the database schema
func SchemaAutoMigrate(db *gorm.DB) {
	log.Infoln("Migrating Database Schema")
	err := db.AutoMigrate(models.Currency{}, models.Pair{}, models.Order{}, models.Peer{}).Error
	if err != nil {
		log.Fatalln(err)
	}
}
