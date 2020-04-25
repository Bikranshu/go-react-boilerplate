package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var dbPrefix = "database"

func DBOpen() *gorm.DB {
	dbDriver := viper.GetString(dbPrefix + ".driver")
	db, err := gorm.Open(dbDriver, buildConnectionOptions(dbDriver))

	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}

	log.Printf("Successfully connected to %s database", dbDriver)
	return db
}

func buildConnectionOptions(dbDriver string) string {
	dbHost := viper.GetString(dbPrefix + ".host")
	dbPort := viper.GetString(dbPrefix + ".port")
	dbUser := viper.GetString(dbPrefix + ".user")
	dbName := viper.GetString(dbPrefix + ".name")
	dbPassword := viper.GetString(dbPrefix + ".password")
	dbSslMode := viper.GetString(dbPrefix + ".sslmode")

	switch dbDriver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	case "postgres":
		return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", dbHost, dbPort, dbUser, dbName, dbPassword, dbSslMode)
	}
	log.Printf("%s driver not supported", dbDriver)
	return ""
}
