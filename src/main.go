package main

import (
	"./migrations"
	"./routes"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"
)

func init() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err.Error())
	}
	os.Setenv("secret", viper.GetString("jwt_secret"))
}

func dbConnect(dbDriver string, dbUser string, dbPassword string, dbHost string, dbPort string, dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(dbDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName))
	return db, err
}

func main() {
	// DB binding
	dbPrefix := "database"
	dbDriver := viper.GetString(dbPrefix + ".driver")
	dbHost := viper.GetString(dbPrefix + ".host")
	dbPort := viper.GetString(dbPrefix + ".port")
	dbUser := viper.GetString(dbPrefix + ".user")
	dbName := viper.GetString(dbPrefix + ".name")
	dbPassword := viper.GetString(dbPrefix + ".password")
	db, err := dbConnect(dbDriver, dbUser, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}
	defer db.Close()
	log.Println("Connected to the database")

	// migrations
	migrations.InitMigration(db)

	// HTTP(s) binding
	serverPrefix := "server"
	serverHost := viper.GetString(serverPrefix + ".host")
	serverPort := os.Getenv("PORT")
	timeout := time.Duration(viper.GetInt("timeout"))

	if serverPort == "" {
		serverPort = viper.GetString(serverPrefix + ".port")
	}

	conn := serverHost + ":" + serverPort

	srv := &http.Server{
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		Addr:         conn,
		Handler:      routes.InitRoute(db),
	}

	log.Printf("Server running on %s", conn)
	log.Fatal(srv.ListenAndServe())
}
