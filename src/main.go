package main

import (
	"./config"
	"./migrations"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func init() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	viper.SetConfigFile("env.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err.Error())
	}
	os.Setenv("secret", viper.GetString("jwt.secret"))
}

// @title Go React Boilerplate
// @version 1.0.0
// @description RESTful API description with Swagger
// @host localhost:3000
// @BasePath /

// @securityDefinitions.bearer BearerAuth
func main() {
	// DB binding
	db := config.DBOpen()
	defer db.Close()

	// Migrations
	migrations.Migrate(db)

	// HTTP(s) binding
	srv := config.RunServer(db)

	log.Fatal(srv.ListenAndServe())
}
