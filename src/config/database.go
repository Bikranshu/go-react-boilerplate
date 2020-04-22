package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func DBOpen() *gorm.DB {
	dbPrefix := "database"
	dbDriver := viper.GetString(dbPrefix + ".driver")
	dbHost := viper.GetString(dbPrefix + ".host")
	dbPort := viper.GetString(dbPrefix + ".port")
	dbUser := viper.GetString(dbPrefix + ".user")
	dbName := viper.GetString(dbPrefix + ".name")
	dbPassword := viper.GetString(dbPrefix + ".password")

	db, err := gorm.Open(dbDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}

	log.Println("Successfully connected to database")

	return db
}
