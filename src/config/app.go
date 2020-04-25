package config

import (
	"../routes"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"
)

// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
func RunServer(db *gorm.DB) *http.Server {
	var serverPrefix = "server"
	var serverHost = viper.GetString(serverPrefix + ".host")
	var serverPort = os.Getenv("PORT")
	if serverPort == "" {
		serverPort = viper.GetString(serverPrefix + ".port")
	}
	var conn = serverHost + ":" + serverPort
	var timeout = viper.GetInt("context.timeout")

	srv := &http.Server{
		ReadTimeout:  time.Duration(timeout) * time.Second,
		WriteTimeout: time.Duration(timeout) * time.Second,
		Addr:         conn,
		Handler:      routes.InitRoute(db),
	}

	log.Printf("Server running on %s", conn)
	return srv
}
